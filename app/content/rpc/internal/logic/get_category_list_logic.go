package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: category
func (l *GetCategoryListLogic) GetCategoryList(in *content.GetCategoryListReq) (*content.GetCategoryListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.CategoryModel.SelectBuilder()
	if in.CreatorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.Name != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name": in.Name,
		})
	}
	if in.QuoteCount != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Gt{
			"quote_count": in.Name,
		})
	}
	categoryListResp, total, err := l.svcCtx.CategoryModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	categoryList := make([]*content.CategoryItem, len(categoryListResp))
	for idx, categoryItem := range categoryListResp {
		categoryList[idx] = genCategoryItem(categoryItem)
	}
	return &content.GetCategoryListResp{
		Data:  categoryList,
		Total: uint64(total),
	}, nil
}
