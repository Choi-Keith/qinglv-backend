package collection

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation_client"
	"qinglv-backend/pkg/snowflake"

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
	_, err = l.svcCtx.OperationRpc.GetCollectionGroupById(l.ctx, &operation_client.GetCollectionGroupByIdReq{
		Id: req.GroupId,
	})
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.OperationRpc.AddCollection(l.ctx, &operation_client.AddCollectionReq{
		Id:        id,
		CreatorId: uint64(userId),
		TargetId:  req.TargetId,
		GroupId:   req.GroupId,
	})
	if err != nil {
		return err
	}
	return nil
}
