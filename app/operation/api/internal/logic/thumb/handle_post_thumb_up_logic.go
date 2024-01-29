package thumb

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation_client"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostThumbUpLogic {
	return &HandlePostThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostThumbUpLogic) HandlePostThumbUp(req *types.HandlePostThumbUpReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	postResp, err := l.svcCtx.ContentRpc.GetPostDetail(l.ctx, &content_client.GetPostDetailReq{
		Id: req.PostId,
	})
	if err != nil {
		return err
	}
	thumbResp, err := l.svcCtx.OperationRpc.GetThumbDetail(l.ctx, &operation_client.GetThumbDetailReq{
		CreatorId: uint64(userId),
		PostId:    req.PostId,
		Type:      1,
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
		if _, err = l.svcCtx.OperationRpc.UpdateThumb(l.ctx, &operation_client.UpdateThumbReq{
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
		if _, err = l.svcCtx.OperationRpc.AddThumb(l.ctx, &operation_client.AddThumbReq{
			Id:        id,
			CreatorId: uint64(userId),
			PostId:    req.PostId,
			Like:      globalKey.ThumbYes,
			Dislike:   globalKey.ThumbNo,
			Type:      1,
		}); err != nil {
			return err
		}
	}
	if _, err = l.svcCtx.ContentRpc.UpdatePost(l.ctx, &content_client.UpdatePostReq{
		Id:           postResp.Post.Id,
		LikeCount:    postResp.Post.LikeCount + uint64(thumbUpCount),
		DislikeCount: postResp.Post.DislikeCount + uint64(thumbDownCount),
		Score:        postResp.Post.Score + uint64(thumbUpCount),
	}); err != nil {
		return err
	}

	return nil
}
