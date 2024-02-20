package thumb

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentThumbDownLogic {
	return &HandleArticleCommentThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentThumbDownLogic) HandleArticleCommentThumbDown(req *types.HandleArticleCommentThumbDownReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	articleCommentResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &operation.GetCommentByIdReq{
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
	if len(thumbResp.Post) != 0 {
		switch {
		case thumbResp.Post[0].Dislike == globalKey.ThumbYes:
			thumbUpCount, thumbDownCount = 0, -1
		case thumbResp.Post[0].Dislike == globalKey.ThumbNo && thumbResp.Post[0].Like == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = 0, 1
		default:
			thumbUpCount, thumbDownCount = -1, 1
			thumbResp.Post[0].Like = globalKey.ThumbNo

		}
		if _, err = l.svcCtx.ThumbRpc.UpdateCommentThumb(l.ctx, &operation.UpdateCommentThumbReq{
			Id:      thumbResp.Post[0].Id,
			Dislike: 1 - thumbResp.Post[0].Dislike,
			Like:    thumbResp.Post[0].Like,
			Type:    2,
		}); err != nil {
			return err
		}
	} else {
		id := snowflake.MustID()
		thumbUpCount, thumbDownCount = 0, 1
		if _, err = l.svcCtx.ThumbRpc.AddCommentThumb(l.ctx, &operation.AddCommentThumbReq{
			Id:          id,
			CreatorId:   uint64(userId),
			ArticleId:   req.ArticleId,
			CommentId:   req.CommentId,
			Like:        globalKey.ThumbNo,
			Dislike:     globalKey.ThumbYes,
			Type:        2,
			CommentType: 1,
		}); err != nil {
			return err
		}
	}
	if _, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &operation.UpdateCommentReq{
		Id:           articleCommentResp.ArticleComment.Id,
		LikeCount:    articleCommentResp.ArticleComment.LikeCount + uint64(thumbUpCount),
		DislikeCount: articleCommentResp.ArticleComment.DislikeCount + uint64(thumbDownCount),
		Type:         2,
	}); err != nil {
		return err
	}

	return nil
}
