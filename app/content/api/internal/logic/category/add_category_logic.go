package category

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.AddCategoryReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.ContentRpc.AddCategory(l.ctx, &content_client.AddCategoryReq{
		Id:          id,
		CreatorId:   uint64(userId),
		Name:        req.Name,
		Image:       req.Image,
		Description: req.Description,
	})
	if err != nil {
		return err
	}
	return nil
}
