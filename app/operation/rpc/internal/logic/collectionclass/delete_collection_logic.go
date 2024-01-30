package collectionclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Collection
func (l *DeleteCollectionLogic) DeleteCollection(in *operation.DeleteCollectionReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.CollectionModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
