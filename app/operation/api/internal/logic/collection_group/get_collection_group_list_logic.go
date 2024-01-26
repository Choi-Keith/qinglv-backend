package collection_group

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/app/operation/rpc/operation_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollectionGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionGroupListLogic {
	return &GetCollectionGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollectionGroupListLogic) GetCollectionGroupList(req *types.GetCollectionGroupListReq) (resp *types.GetCollectionGropListResp, err error) {
	// todo: add your logic here and delete this line
	groupListResp, err := l.svcCtx.OperationRpc.GetCollectionGroupList(l.ctx, &operation_client.GetCollectionGroupListReq{
		Name:       req.Name,
		Visibility: req.Visibility,
		BizType:    req.BizType,
		CreatorId:  req.CreatorId,
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	groupList := make([]types.CollectionGroupItem, len(groupListResp.Data))
	for idx, groupItem := range groupListResp.Data {
		_ = copier.Copy(&groupList[idx], groupItem)
	}
	isEnd := false
	total := (uint64(req.PageNum)-1)*uint64(req.PageSize) + uint64(req.PageSize)
	if groupListResp.Total <= total {
		isEnd = true
	}
	return &types.GetCollectionGropListResp{
		List:  groupList,
		Total: groupListResp.Total,
		IsEnd: isEnd,
	}, nil
}
