package logic

import (
	"context"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/collection"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Collection
func (l *AddCollectionLogic) AddCollection(in *operation.AddCollectionReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.CollectionModel.Insert(l.ctx, nil, &collection.Collection{
		Id:        in.Id,
		TargetId:  in.TargetId,
		GroupId:   in.GroupId,
		CreatorId: in.CreatorId,
		DeletedAt: time.Now(),
		Version:   1,
	})
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
