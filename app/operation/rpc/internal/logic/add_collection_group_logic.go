package logic

import (
	"context"
	"time"

	"qinglv-backend/app/operation/rpc/internal/model/collection_group"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectionGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionGroupLogic {
	return &AddCollectionGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CollectionGroup
func (l *AddCollectionGroupLogic) AddCollectionGroup(in *operation.AddCollectionGroupReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.CollectionGroupModel.Insert(l.ctx, nil, &collection_group.CollectionGroup{
		Id:         in.Id,
		CreatorId:  in.CreatorId,
		Name:       in.Name,
		Visibility: uint64(in.Visibility),
		BizType:    uint64(in.BizType),
		DeletedAt:  time.Now(),
		Version:    1,
	})
	if err != nil {
		return nil, err
	}
	return &operation.OkResp{}, nil
}
