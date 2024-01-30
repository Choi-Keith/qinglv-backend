package collection_group

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCollectionGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCollectionGroupLogic {
	return &UpdateCollectionGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCollectionGroupLogic) UpdateCollectionGroup(req *types.UpdateCollectionGroup) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	groupListResp, err := l.svcCtx.CollectionRpc.CheckExistByName(l.ctx, &operation.CheckExistByNameReq{
		Name:      req.Name,
		CreatorId: uint64(userId),
		BizType:   req.BizType,
	})
	if err == nil && len(groupListResp.Data) > 0 {
		return errors.New("该收藏夹已存在")
	}
	_, err = l.svcCtx.CollectionRpc.UpdateCollectionGroup(l.ctx, &operation.UpdateCollectionGroupReq{
		Id:         req.Id,
		BizType:    req.BizType,
		Name:       req.Name,
		Visibility: req.Visibility,
	})
	if err != nil {
		return err
	}
	return nil
}
