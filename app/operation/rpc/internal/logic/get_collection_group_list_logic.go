package logic

import (
	"context"
	"fmt"

	operationModel "qinglv-backend/app/operation/rpc/internal/model/collection_group"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionGroupListLogic {
	return &GetCollectionGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CollectionGroup
func (l *GetCollectionGroupListLogic) GetCollectionGroupList(in *operation.GetCollectionGroupListReq) (*operation.GetCollectionGroupListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.CollectionGroupModel.SelectBuilder()
	if in.BizType != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"biz_type": in.BizType,
		})
	}
	if in.CreatorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.Visibility != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.Name != "" {
		whereBuilder = whereBuilder.Where("name LIKE ?", fmt.Sprint("%", in.Name, "%"))
	}
	groupListResp, total, err := l.svcCtx.CollectionGroupModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	groupList := make([]*operation.CollectionGroupItem, len(groupListResp))
	for idx, groupItem := range groupListResp {
		groupList[idx] = genCollectionGroupItem(groupItem)
	}

	return &operation.GetCollectionGroupListResp{
		Data:  groupList,
		Total: uint64(total),
	}, nil
}

func genCollectionGroupItem(groupItem *operationModel.CollectionGroup) *operation.CollectionGroupItem {
	return &operation.CollectionGroupItem{
		Id:         groupItem.Id,
		CreatorId:  groupItem.CreatorId,
		Name:       groupItem.Name,
		Visibility: int32(groupItem.Visibility),
		BizType:    int32(groupItem.BizType),
		CreatedAt:  uint64(groupItem.CreatedAt.Unix() * 1000),
		UpdatedAt:  uint64(groupItem.UpdatedAt.Unix() * 1000),
	}
}
