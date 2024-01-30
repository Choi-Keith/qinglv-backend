package collectionclasslogic

import (
	"context"

	collectionModel "qinglv-backend/app/operation/rpc/internal/model/collection"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionByIdLogic {
	return &GetCollectionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Collection
func (l *GetCollectionByIdLogic) GetCollectionById(in *operation.GetCollectionByIdReq) (*operation.GetCollectionByIdResp, error) {
	// todo: add your logic here and delete this line
	collectionResp, err := l.svcCtx.CollectionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	collectionItem := genCollectionItem(collectionResp)
	return &operation.GetCollectionByIdResp{
		Collection: collectionItem,
	}, nil
}

func genCollectionItem(collectionItem *collectionModel.Collection) *operation.CollectionItem {
	return &operation.CollectionItem{
		Id:        collectionItem.Id,
		GroupId:   collectionItem.GroupId,
		TargetId:  collectionItem.TargetId,
		CreatorId: collectionItem.CreatorId,
		CreatedAt: uint64(collectionItem.CreatedAt.Unix() * 1000),
		UpdatedAt: uint64(collectionItem.UpdatedAt.Unix() * 1000),
	}
}
