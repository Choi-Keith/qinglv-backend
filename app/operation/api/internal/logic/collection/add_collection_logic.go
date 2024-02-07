package collection

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"
	"qinglv-backend/pkg/snowflake"
	"qinglv-backend/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollectionLogic) AddCollection(req *types.AddCollectionReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	groupResp, err := l.svcCtx.CollectionRpc.GetCollectionGroupById(l.ctx, &operation.GetCollectionGroupByIdReq{
		Id: req.GroupId,
	})
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	if _, err = l.svcCtx.CollectionRpc.AddCollection(l.ctx, &operation.AddCollectionReq{
		Id:        id,
		CreatorId: uint64(userId),
		TargetId:  req.TargetId,
		GroupId:   req.GroupId,
	}); err != nil {
		return err
	}
	if groupResp.CollectionGroup.BizType == 1 {
		postResp, err := l.svcCtx.PostRpc.GetPostDetail(l.ctx, &content.GetPostDetailReq{
			Id: req.TargetId,
		})
		if err != nil {
			return err
		}
		score := utils.HandleScore(postResp.Post.CreatedAt, 2, 1.5)
		if _, err := l.svcCtx.PostRpc.UpdatePost(l.ctx, &content.UpdatePostReq{
			Id:              req.TargetId,
			CollectionCount: postResp.Post.CollectionCount + 1,
			Score:           postResp.Post.Score + score,
		}); err != nil {
			return err
		}
	}

	return nil
}
