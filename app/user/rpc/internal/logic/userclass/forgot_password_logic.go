package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgotPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgotPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgotPasswordLogic {
	return &ForgotPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ForgotPasswordLogic) ForgotPassword(in *user.ForgotPasswordReq) (*user.OkResp, error) {
	// todo: add your logic here and delete this line
	if err := sendAndSaveForgotCode(l.svcCtx, in.UserId, in.Email); err != nil {
		return nil, err
	}
	return &user.OkResp{}, nil
}
