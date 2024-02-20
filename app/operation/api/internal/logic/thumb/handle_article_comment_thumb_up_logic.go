package thumb

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/client/thumbclass"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentThumbUpLogic {
	return &HandleArticleCommentThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentThumbUpLogic) HandleArticleCommentThumbUp(req *types.HandleArticleCommentThumbUpReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	creatorName := l.ctx.Value("nickname").(string)
	if err != nil {
		return err
	}
	commentResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &operation.GetCommentByIdReq{
		Id:   req.CommentId,
		Type: 2,
	})
	if err != nil {
		return err
	}
	thumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
		CreatorId: uint64(userId),
		ArticleId: req.ArticleId,
		CommentId: req.CommentId,
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
		if _, err = l.svcCtx.ThumbRpc.UpdateCommentThumb(l.ctx, &thumbclass.UpdateCommentThumbReq{
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
		if _, err = l.svcCtx.ThumbRpc.AddCommentThumb(l.ctx, &thumbclass.AddCommentThumbReq{
			Id:          id,
			CreatorId:   uint64(userId),
			ArticleId:   req.ArticleId,
			CommentId:   req.CommentId,
			Like:        globalKey.ThumbYes,
			Dislike:     globalKey.ThumbNo,
			CreatorName: creatorName,
			Type:        2,
			CommentType: 1,
		}); err != nil {
			return err
		}
	}
	if _, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &operation.UpdateCommentReq{
		Id:           commentResp.ArticleComment.Id,
		LikeCount:    commentResp.ArticleComment.LikeCount + uint64(thumbUpCount),
		DislikeCount: commentResp.ArticleComment.DislikeCount + uint64(thumbDownCount),
		Type:         2,
	}); err != nil {
		return err
	}

	return nil
}
