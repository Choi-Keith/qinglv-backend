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

type HandlePostCommentThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentThumbDownLogic {
	return &HandlePostCommentThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentThumbDownLogic) HandlePostCommentThumbDown(req *types.HandlePostCommentThumbDownReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	postCommentResp, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &operation.GetCommentByIdReq{
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
			Type:    1,
		}); err != nil {
			return err
		}
	} else {
		id := snowflake.MustID()
		thumbUpCount, thumbDownCount = 0, 1
		if _, err = l.svcCtx.ThumbRpc.AddCommentThumb(l.ctx, &operation.AddCommentThumbReq{
			Id:          id,
			CreatorId:   uint64(userId),
			PostId:      req.PostId,
			CommentId:   req.CommentId,
			Like:        globalKey.ThumbNo,
			Dislike:     globalKey.ThumbYes,
			Type:        1,
			CommentType: 1,
		}); err != nil {
			return err
		}
	}
	if _, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &operation.UpdateCommentReq{
		Id:           postCommentResp.PostComment.Id,
		LikeCount:    postCommentResp.PostComment.LikeCount + uint64(thumbUpCount),
		DislikeCount: postCommentResp.PostComment.DislikeCount + uint64(thumbDownCount),
		Type:         1,
	}); err != nil {
		return err
	}

	return nil
}
