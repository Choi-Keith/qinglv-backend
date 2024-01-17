package logic

import (
	"context"
	"database/sql"
	"time"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/category"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: category
func (l *AddCategoryLogic) AddCategory(in *content.AddCategoryReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.CategoryModel.Insert(l.ctx, nil, &category.Category{
		Id:          in.Id,
		CreatorId:   in.CreatorId,
		Name:        in.Name,
		Description: sql.NullString{String: in.Description, Valid: true},
		Image:       in.Image,
		QuoteCount:  0,
		Version:     1,
		DeletedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
