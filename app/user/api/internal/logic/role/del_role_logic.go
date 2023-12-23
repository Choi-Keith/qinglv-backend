package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"qinglv-backend/app/user/api/internal/svc"
)

type DelRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRoleLogic {
	return &DelRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRoleLogic) DelRole() error {
	// todo: add your logic here and delete this line

	return nil
}
