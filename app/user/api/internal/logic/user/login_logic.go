package user

import (
	"context"
	"errors"
	"fmt"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/jwtx"
	"qinglv-backend/pkg/password"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	result, _ := l.svcCtx.UserRpc.GetUserInfoByParams(l.ctx, &user_client.GetUserInfoByParamsReq{
		Email: req.Email,
	})
	if result == nil {
		return nil, errors.New(fmt.Sprintf("不存在邮箱为%s的用户", req.Email))
	}
	if result.User.MailStatus == 0 {
		return nil, errors.New("该邮箱不可用，请重新注册")
	}
	if result.User.Status != 1 {
		return nil, errors.New("该用户已被注销")
	}
	verifyCaptchaResp, err := l.svcCtx.UserRpc.VerifyCaptcha(l.ctx, &user_client.VerifyCaptchaReq{
		Key:     req.CaptchaCode,
		Captcha: req.CaptchaValue,
	})
	if err != nil {
		return nil, err
	}
	if !verifyCaptchaResp.IsCorrect {
		return nil, errors.New("验证码错误")
	}
	isEqual := password.VerifyPassword(req.Password, result.User.Password)
	if !isEqual {
		return nil, errors.New("密码错误，请重新输入")
	}
	secretKey := l.svcCtx.Config.JWTAuth.AccessSecret
	seconds := l.svcCtx.Config.JWTAuth.AccessExpire
	token, err := jwtx.NewJwtToken(secretKey, seconds, jwtx.WithOption("userId", result.User.Id), jwtx.WithOption("email", result.User.Email))
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  token,
		ExpireAt:     uint64(seconds),
		RefreshAfter: uint64(seconds / 2),
	}, nil
}
