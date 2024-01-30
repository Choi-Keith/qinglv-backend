package collectionclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCollectionGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCollectionGroupLogic {
	return &UpdateCollectionGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Collection
func (l *UpdateCollectionGroupLogic) UpdateCollectionGroup(in *operation.UpdateCollectionGroupReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	group, err := l.svcCtx.CollectionGroupModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.CollectionGroupModel.UpdateWithVersion(l.ctx, nil, group)
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
