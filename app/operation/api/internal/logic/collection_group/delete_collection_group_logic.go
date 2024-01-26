package collection_group

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionGroupLogic {
	return &DeleteCollectionGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectionGroupLogic) DeleteCollectionGroup(req *types.DeleteCollectionGroup) error {
	// todo: add your logic here and delete this line
	return nil
}
