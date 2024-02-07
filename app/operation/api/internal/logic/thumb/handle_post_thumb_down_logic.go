package thumb

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/client/thumbclass"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostThumbDownLogic {
	return &HandlePostThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostThumbDownLogic) HandlePostThumbDown(req *types.HandlePostThumbDownReq) error {
	// todo: add your logic here and delete this line
	var (
		thumbUpCount   = 0
		thumbDownCount = 0
	)
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
		Id: req.PostId,
	})
	if err != nil {
		return err
	}
	thumbResp, err := l.svcCtx.ThumbRpc.GetThumbDetail(l.ctx, &thumbclass.GetThumbDetailReq{
		CreatorId: uint64(userId),
		PostId:    req.PostId,
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
		if _, err = l.svcCtx.ThumbRpc.UpdateThumb(l.ctx, &thumbclass.UpdateThumbReq{
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
		if _, err = l.svcCtx.ThumbRpc.AddThumb(l.ctx, &thumbclass.AddThumbReq{
			Id:        id,
			CreatorId: uint64(userId),
			PostId:    req.PostId,
			Like:      globalKey.ThumbNo,
			Dislike:   globalKey.ThumbYes,
			Type:      1,
		}); err != nil {
			return err
		}
	}
	var score float64
	if thumbUpCount > 0 {
		score = postResp.Post.Score - utils.HandleScore(postResp.Post.CreatedAt, 1, 1.5)
	} else {
		score = postResp.Post.Score + utils.HandleScore(postResp.Post.CreatedAt, 1, 1.5)
	}
	status := 1
	if postResp.Post.LikeCount+uint64(thumbUpCount) < (postResp.Post.DislikeCount+uint64(thumbDownCount))*2 && postResp.Post.DislikeCount+uint64(thumbDownCount) > 5 {
		status = 2
	}
	if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
		Id:           postResp.Post.Id,
		LikeCount:    postResp.Post.LikeCount + uint64(thumbUpCount),
		DislikeCount: postResp.Post.DislikeCount + uint64(thumbDownCount),
		Score:        score,
		Status:       int32(status),
	}); err != nil {
		return err
	}

	return nil
}
