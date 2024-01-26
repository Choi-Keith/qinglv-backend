package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckExistByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckExistByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckExistByNameLogic {
	return &CheckExistByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CollectionGroup
func (l *CheckExistByNameLogic) CheckExistByName(in *operation.CheckExistByNameReq) (*operation.CheckExistByNameResp, error) {
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
	if in.Name != "" {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"name": in.Name,
		})
	}
	groupListResp, err := l.svcCtx.CollectionGroupModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	groupList := make([]*operation.CollectionGroupItem, len(groupListResp))
	for idx, groupItem := range groupListResp {
		groupList[idx] = genCollectionGroupItem(groupItem)
	}
	return &operation.CheckExistByNameResp{
		Data: groupList,
	}, nil
}
