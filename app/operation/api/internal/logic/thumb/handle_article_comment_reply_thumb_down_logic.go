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

type HandleArticleCommentReplyThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentReplyThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentReplyThumbDownLogic {
	return &HandleArticleCommentReplyThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentReplyThumbDownLogic) HandleArticleCommentReplyThumbDown(req *types.HandleArticleCommentReplyThumbDownReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	articleCommentResp, err := l.svcCtx.CommentRpc.GetCommentReplyById(l.ctx, &operation.GetCommentReplyByIdReq{
		Id:   req.ReplyId,
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
		ReplyId:   req.ReplyId,
	})
	if err != nil {
		return err
	}
	if len(thumbResp.Article) != 0 {
		switch {
		case thumbResp.Article[0].Dislike == globalKey.ThumbYes:
			thumbUpCount, thumbDownCount = 0, -1
		case thumbResp.Article[0].Dislike == globalKey.ThumbNo && thumbResp.Article[0].Like == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = 0, 1
		default:
			thumbUpCount, thumbDownCount = -1, 1
			thumbResp.Article[0].Like = globalKey.ThumbNo

		}
		if _, err = l.svcCtx.ThumbRpc.UpdateCommentThumb(l.ctx, &operation.UpdateCommentThumbReq{
			Id:      thumbResp.Article[0].Id,
			Dislike: 1 - thumbResp.Article[0].Dislike,
			Like:    thumbResp.Article[0].Like,
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
			ReplyId:     req.ReplyId,
			Like:        globalKey.ThumbNo,
			Dislike:     globalKey.ThumbYes,
			Type:        2,
			CommentType: 2,
		}); err != nil {
			return err
		}
	}
	if _, err = l.svcCtx.CommentRpc.UpdateCommentReply(l.ctx, &operation.UpdateCommentReplyReq{
		Id:           articleCommentResp.ArticleCommentReply.Id,
		LikeCount:    articleCommentResp.ArticleCommentReply.LikeCount + uint64(thumbUpCount),
		DislikeCount: articleCommentResp.ArticleCommentReply.DislikeCount + uint64(thumbDownCount),
		Type:         2,
	}); err != nil {
		return err
	}

	return nil
}
