package category

import (
	"context"
	"strings"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.UpdateCategoryReq) error {
	// todo: add your logic here and delete this line
	var oldImage string
	if req.Image != "" {
		categoryResp, err := l.svcCtx.CategoryRpc.GetCategoryDetail(l.ctx, &content.GetCategoryDetailReq{
			Id: req.Id,
		})
		if err != nil {
			return err
		}
		oldImage = categoryResp.Category.Image
	}
	_, err := l.svcCtx.CategoryRpc.UpdateCategory(l.ctx, &content.UpdateCategoryReq{
		Id:          req.Id,
		Image:       req.Image,
		Description: req.Description,
	})
	if err != nil {
		return err
	}
	if req.Image != oldImage && oldImage != "" {
		name, _ := strings.CutPrefix(oldImage, l.svcCtx.Config.Cos.Endpoint)
		_, err := l.svcCtx.CosClient.Object.Delete(context.Background(), name)
		if err != nil {
			return err
		}
	}
	return nil
}
