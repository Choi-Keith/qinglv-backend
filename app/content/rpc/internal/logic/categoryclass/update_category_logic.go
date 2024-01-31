package categoryclasslogic

import (
	"context"
	"database/sql"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: category
func (l *UpdateCategoryLogic) UpdateCategory(in *content.UpdateCategoryReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	categoryResp, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	categoryResp.Description = sql.NullString{String: in.Description, Valid: true}
	categoryResp.Image = in.Image
	categoryResp.QuoteCount = in.QuoteCount
	err = l.svcCtx.CategoryModel.UpdateWithVersion(l.ctx, nil, categoryResp)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
