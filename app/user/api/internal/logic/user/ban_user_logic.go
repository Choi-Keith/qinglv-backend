package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"qinglv-backend/app/user/api/internal/svc"
)

type BanUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBanUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BanUserLogic {
	return &BanUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BanUserLogic) BanUser() error {
	// todo: add your logic here and delete this line

	return nil
}
