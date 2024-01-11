package user

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *UpdatePasswordLogic) UpdatePassword(in *user.UpdatePasswordReq) (*user.UpdatePasswordResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdatePasswordResp{}, nil
}
