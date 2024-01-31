package categoryclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	categoryModel "qinglv-backend/app/content/rpc/internal/model/category"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryDetailLogic {
	return &GetCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: category
func (l *GetCategoryDetailLogic) GetCategoryDetail(in *content.GetCategoryDetailReq) (*content.GetCategoryDetailResp, error) {
	// todo: add your logic here and delete this line
	categoryResp, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	categoryItem := genCategoryItem(categoryResp)

	return &content.GetCategoryDetailResp{
		Category: categoryItem,
	}, nil
}

func genCategoryItem(categoryItem *categoryModel.Category) *content.CategoryItem {
	return &content.CategoryItem{
		Id:         categoryItem.Id,
		CreatorId:  categoryItem.CreatorId,
		Name:       categoryItem.Name,
		Image:      categoryItem.Image,
		QuoteCount: categoryItem.QuoteCount,
		CreatedAt:  uint64(categoryItem.CreatedAt.Unix() * 1000),
		UpdatedAt:  uint64(categoryItem.UpdatedAt.Unix() * 1000),
	}
}
