package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionGroupByIdLogic {
	return &GetCollectionGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CollectionGroup
func (l *GetCollectionGroupByIdLogic) GetCollectionGroupById(in *operation.GetCollectionGroupByIdReq) (*operation.GetCollectionGroupByIdResp, error) {
	// todo: add your logic here and delete this line
	collectionGroup, err := l.svcCtx.CollectionGroupModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	collectionGroupItem := genCollectionGroupItem(collectionGroup)
	return &operation.GetCollectionGroupByIdResp{
		CollectionGroup: collectionGroupItem,
	}, nil
}
