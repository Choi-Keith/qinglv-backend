package user

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/password"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.PasswordReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: uint64(userId),
	})
	if err != nil {
		return err
	}
	isEqual := password.VerifyPassword(userResp.User.Password, req.OldPassword)
	if !isEqual {
		return errors.New("原密码错误，请重新输入")
	}
	newPassword, err := password.EncryptPassword(req.NewPassword)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.UserRpc.UpdatePassword(l.ctx, &user.UpdatePasswordReq{
		UserId:      uint64(userId),
		NewPassword: newPassword,
	})
	if err != nil {
		return err
	}
	return nil
}
