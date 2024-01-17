package category

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(req *types.GetCategoryByIdReq) (resp *types.GetCategoryByIdResp, err error) {
	// todo: add your logic here and delete this line
	categoryResp, err := l.svcCtx.ContentRpc.GetCategoryDetail(l.ctx, &content_client.GetCategoryDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
		UserId: categoryResp.Category.CreatorId,
	})
	if err != nil {
		return nil, err
	}
	var categoryItem types.CategoryItem
	_ = copier.Copy(&categoryItem, categoryResp.Category)
	_ = copier.Copy(&categoryItem.Creator, userResp.User)
	return &types.GetCategoryByIdResp{
		Category: categoryItem,
	}, nil
}
