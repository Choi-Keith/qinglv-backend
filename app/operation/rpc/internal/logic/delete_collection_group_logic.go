package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionGroupLogic {
	return &DeleteCollectionGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CollectionGroup
func (l *DeleteCollectionGroupLogic) DeleteCollectionGroup(in *operation.DeleteCollectionGroupReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	return &operation.OkResp{}, nil
}
