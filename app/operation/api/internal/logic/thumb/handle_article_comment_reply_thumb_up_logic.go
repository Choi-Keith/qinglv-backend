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

type HandleArticleCommentReplyThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentReplyThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentReplyThumbUpLogic {
	return &HandleArticleCommentReplyThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentReplyThumbUpLogic) HandleArticleCommentReplyThumbUp(req *types.HandleArticleCommentReplyThumbUpReq) error {
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
		ReplyId:   req.ReplyId,
		Type:      2,
	})
	if err != nil {
		return err
	}
	logx.Debugf("thumbResp: %+v\n", thumbResp)
	if len(thumbResp.Post) != 0 {
		switch {
		case thumbResp.Post[0].Like == globalKey.ThumbYes && thumbResp.Post[0].Dislike == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = -1, 0
		case thumbResp.Post[0].Like == globalKey.ThumbNo && thumbResp.Post[0].Dislike == globalKey.ThumbNo:
			thumbUpCount, thumbDownCount = 1, 0
		default:
			thumbUpCount, thumbDownCount = 1, -1
			thumbResp.Post[0].Dislike = globalKey.ThumbNo
		}
		if _, err = l.svcCtx.ThumbRpc.UpdateCommentThumb(l.ctx, &thumbclass.UpdateCommentThumbReq{
			Id:      thumbResp.Post[0].Id,
			Like:    1 - thumbResp.Post[0].Like,
			Dislike: thumbResp.Post[0].Dislike,
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
			ReplyId:     req.ReplyId,
			Like:        globalKey.ThumbYes,
			Dislike:     globalKey.ThumbNo,
			CreatorName: creatorName,
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
