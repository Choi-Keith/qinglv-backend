package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"qinglv-backend/app/user/api/internal/svc"
)

type GetRoleByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIDLogic {
	return &GetRoleByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByIDLogic) GetRoleByID() error {
	// todo: add your logic here and delete this line

	return nil
}
