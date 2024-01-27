package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionListLogic {
	return &GetCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Collection
func (l *GetCollectionListLogic) GetCollectionList(in *operation.GetCollectionListReq) (*operation.GetCollectionListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.CollectionModel.SelectBuilder()
	if in.GroupId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"group_id": in.GroupId,
		})
	}
	collectionListResp, total, err := l.svcCtx.CollectionModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	collectionList := make([]*operation.CollectionItem, len(collectionListResp))
	for idx, collectionItem := range collectionListResp {
		collectionList[idx] = genCollectionItem(collectionItem)
	}
	return &operation.GetCollectionListResp{
		Data:  collectionList,
		Total: uint64(total),
	}, nil
}
