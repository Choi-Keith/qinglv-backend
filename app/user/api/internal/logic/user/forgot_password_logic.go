package user

import (
	"context"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgotPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgotPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgotPasswordLogic {
	return &ForgotPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgotPasswordLogic) ForgotPassword(req *types.ForgotPasswordReq) error {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpc.CheckEmailExist(l.ctx, &user.CheckEmailExistReq{
		Email: req.Email,
	})
	if err != nil {
		return err
	}
	if !userResp.IsExist {
		return errors.New("邮箱不存在")
	}
	if _, err := l.svcCtx.UserRpc.ForgotPassword(l.ctx, &user.ForgotPasswordReq{
		UserId: userResp.User.Id,
		Email:  req.Email,
	}); err != nil {
		return err
	}
	return nil
}
