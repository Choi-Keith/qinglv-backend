package user

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/schema"
	"qinglv-backend/pkg/password"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) error {
	// todo: add your logic here and delete this line
	verifyRegisterCodeResp, err := l.svcCtx.EmailRpc.VerifyForgotEemailCode(l.ctx, &user.VerifyForgotEemailCodeReq{
		Code: req.Code,
	})
	if err != nil {
		return err
	}
	codeContent := new(schema.EmailContent)
	if err := json.Unmarshal([]byte(verifyRegisterCodeResp.CodeContent), codeContent); err != nil {
		return err
	}
	newPassword, err := password.EncryptPassword(req.Password)
	if err != nil {
		return err
	}
	if _, err := l.svcCtx.UserRpc.UpdatePassword(l.ctx, &user.UpdatePasswordReq{
		UserId:      codeContent.UserId,
		NewPassword: newPassword,
	}); err != nil {
		return err
	}
	return nil
}
