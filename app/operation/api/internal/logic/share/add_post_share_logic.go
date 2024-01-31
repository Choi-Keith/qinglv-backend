package share

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostShareLogic {
	return &AddPostShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostShareLogic) AddPostShare(req *types.AddPostShareReq) error {
	// todo: add your logic here and delete this line
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
	id := snowflake.MustID()
	if _, err = l.svcCtx.ShareRpc.AddPostShare(l.ctx, &operation.AddPostShareReq{
		Id:        id,
		CreatorId: uint64(userId),
		PostId:    req.PostId,
		Type:      req.Type,
	}); err != nil {
		return err
	}

	if _, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
		Id:              req.PostId,
		ShareCount:      postResp.Post.ShareCount + 1,
		Score:           postResp.Post.Score + 1,
		CollectionCount: postResp.Post.CollectionCount,
	}); err != nil {
		return err
	}
	return nil
}
