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

type HandlePostCommentThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentThumbUpLogic {
	return &HandlePostCommentThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentThumbUpLogic) HandlePostCommentThumbUp(req *types.HandlePostCommentThumbUpReq) error {
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
		Type: 1,
	})
	if err != nil {
		return err
	}
	thumbResp, err := l.svcCtx.ThumbRpc.GetCommentThumbDetail(l.ctx, &operation.GetCommentThumbDetailReq{
		CreatorId: uint64(userId),
		PostId:    req.PostId,
		CommentId: req.CommentId,
		Type:      1,
	})
	if err != nil {
		return err
	}
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
			Type:    1,
		}); err != nil {
			return err
		}

	} else {
		id := snowflake.MustID()
		thumbUpCount, thumbDownCount = 1, 0
		if _, err = l.svcCtx.ThumbRpc.AddCommentThumb(l.ctx, &thumbclass.AddCommentThumbReq{
			Id:          id,
			CreatorId:   uint64(userId),
			PostId:      req.PostId,
			CommentId:   req.CommentId,
			Like:        globalKey.ThumbYes,
			Dislike:     globalKey.ThumbNo,
			CreatorName: creatorName,
			Type:        1,
			CommentType: 1,
		}); err != nil {
			return err
		}
	}
	logx.Debugf("HandlePostCommentThumbUp commentResp: %+v\n", commentResp)
	logx.Debugf("HandlePostCommentThumbUp commentResp.PostComment: %+v\n", commentResp.PostComment)
	if _, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &operation.UpdateCommentReq{
		Id:           commentResp.PostComment.Id,
		LikeCount:    commentResp.PostComment.LikeCount + uint64(thumbUpCount),
		DislikeCount: commentResp.PostComment.DislikeCount + uint64(thumbDownCount),
		Score:        commentResp.PostComment.Score + uint64(thumbUpCount),
		Type:         1,
	}); err != nil {
		return err
	}

	return nil
}
