package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"qinglv-backend/app/user/api/internal/svc"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser() error {
	// todo: add your logic here and delete this line

	return nil
}
