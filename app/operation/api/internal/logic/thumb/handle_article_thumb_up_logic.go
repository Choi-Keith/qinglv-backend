package thumb

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/client/thumbclass"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleThumbUpLogic {
	return &HandleArticleThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleThumbUpLogic) HandleArticleThumbUp(req *types.HandleArticleThumbUpReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	articleResp, err := l.svcCtx.ArticleRpc.GetArticleDetail(l.ctx, &content.GetArticleDetailReq{
		Id: req.ArticleId,
	})
	if err != nil {
		return err
	}
	articleContentResp, err := l.svcCtx.ArticleRpc.GetArticleContentByArticleId(l.ctx, &content.GetArticleContentDetailReq{
		Id: articleResp.Article.Id,
	})
	if err != nil {
		return err
	}
	thumbResp, err := l.svcCtx.ThumbRpc.GetThumbDetail(l.ctx, &thumbclass.GetThumbDetailReq{
		CreatorId: uint64(userId),
		ArticleId: req.ArticleId,
		Type:      2,
	})
	if err != nil {
		return err
	}
	if len(thumbResp.Article) != 0 {
		switch {
		case thumbResp.Article[0].Like == globalKey.ThumbYes && thumbResp.Article[0].Dislike == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = -1, 0
		case thumbResp.Article[0].Like == globalKey.ThumbNo && thumbResp.Article[0].Dislike == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = 1, 0
		default:
			thumbUpCount, thumbDownCount = 1, -1
			thumbResp.Article[0].Dislike = globalKey.ThumbNo
		}
		if _, err = l.svcCtx.ThumbRpc.UpdateThumb(l.ctx, &thumbclass.UpdateThumbReq{
			Id:      thumbResp.Article[0].Id,
			Like:    1 - thumbResp.Article[0].Like,
			Dislike: thumbResp.Article[0].Dislike,
			Type:    2,
		}); err != nil {
			return err
		}
	} else {
		id := snowflake.MustID()
		thumbUpCount, thumbDownCount = 1, 0
		if _, err = l.svcCtx.ThumbRpc.AddThumb(l.ctx, &thumbclass.AddThumbReq{
			Id:        id,
			CreatorId: uint64(userId),
			ArticleId: req.ArticleId,
			Like:      globalKey.ThumbYes,
			Dislike:   globalKey.ThumbNo,
			Type:      2,
		}); err != nil {
			return err
		}
	}
	var score float64
	if thumbUpCount > 0 {
		score = articleResp.Article.Score + utils.HandleScore(articleResp.Article.CreatedAt, 2, 1.5)
	} else {
		score = articleResp.Article.Score - utils.HandleScore(articleResp.Article.CreatedAt, 2, 1.5)
	}
	if _, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
		Id:           articleResp.Article.Id,
		LikeCount:    articleResp.Article.LikeCount + uint64(thumbUpCount),
		DislikeCount: articleResp.Article.DislikeCount + uint64(thumbDownCount),
		Score:        score,
	}); err != nil {
		return err
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    articleResp.Article.CreatorId,
		Score: int32(thumbUpCount) * 2,
		Op:    "add",
	}); err != nil {
		return err
	}
	go l.svcCtx.NotificationRpc.AddNotification(l.ctx, &user.AddNotificationReq{
		Id:             snowflake.MustID(),
		Type:           2,
		ActionType:     1,
		SenderUserId:   uint64(userId),
		ReceiverUserId: articleResp.Article.CreatorId,
		BizType:        2,
		TargetId:       articleResp.Article.Id,
		TargetTitle:    articleContentResp.ArticleContent.Title,
	})
	return nil
}
