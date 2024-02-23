package collectionclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionAllLogic {
	return &GetCollectionAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCollectionAllLogic) GetCollectionAll(in *operation.GetCollectionAllReq) (*operation.GetCollectionAllResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.CollectionModel.SelectBuilder()
	if in.GroupId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"group_id": in.GroupId,
		})
	}
	if in.TargetId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"target_id": in.TargetId,
		})
	}
	collectionListResp, err := l.svcCtx.CollectionModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	collectionList := make([]*operation.CollectionItem, len(collectionListResp))
	for idx, collectionItem := range collectionListResp {
		collectionList[idx] = genCollectionItem(collectionItem)
	}
	return &operation.GetCollectionAllResp{
		Data: collectionList,
	}, nil
}
