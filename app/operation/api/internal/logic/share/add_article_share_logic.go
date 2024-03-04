package share

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleShareLogic {
	return &AddArticleShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleShareLogic) AddArticleShare(req *types.AddArticleShareReq) error {
	// todo: add your logic here and delete this line
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
		Id: req.ArticleId,
	})
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	if _, err = l.svcCtx.ShareRpc.AddArticleShare(l.ctx, &operation.AddArticleShareReq{
		Id:        id,
		CreatorId: uint64(userId),
		ArticleId: req.ArticleId,
		Type:      req.Type,
	}); err != nil {
		return err
	}
	score := utils.HandleScore(articleResp.Article.CreatedAt, 4, 1.5)
	if _, err = l.svcCtx.ArticleRpc.UpdateArticle(l.ctx, &content.UpdateArticleReq{
		Id:              req.ArticleId,
		ShareCount:      articleResp.Article.ShareCount + 1,
		Score:           articleResp.Article.Score + score,
		CollectionCount: articleResp.Article.CollectionCount,
	}); err != nil {
		return err
	}
	if _, err := l.svcCtx.UserRpc.UpdateUserScoreLevel(l.ctx, &user.UpdateUserScoreLevelReq{
		Id:    articleResp.Article.CreatorId,
		Score: 2,
		Op:    "add",
	}); err != nil {
		return err
	}
	go l.svcCtx.NotificationRpc.AddNotification(l.ctx, &user.AddNotificationReq{
		Id:             snowflake.MustID(),
		Type:           2,
		ActionType:     3,
		SenderUserId:   uint64(userId),
		ReceiverUserId: articleResp.Article.CreatorId,
		BizType:        2,
		TargetId:       articleResp.Article.Id,
		TargetTitle:    articleContentResp.ArticleContent.Title,
	})
	return nil
}
