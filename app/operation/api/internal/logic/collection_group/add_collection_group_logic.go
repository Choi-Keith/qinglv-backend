package collection_group

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation_client"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionGroupLogic {
	return &AddCollectionGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollectionGroupLogic) AddCollectionGroup(req *types.AddCollectionGroup) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.OperationRpc.AddCollectionGroup(l.ctx, &operation_client.AddCollectionGroupReq{
		Id:         id,
		CreatorId:  uint64(userId),
		Name:       req.Name,
		Visibility: req.Visibility,
		BizType:    req.BizType,
	})
	if err != nil {
		return err
	}
	return nil
}
