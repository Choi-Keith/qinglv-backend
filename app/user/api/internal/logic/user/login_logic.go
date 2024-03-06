package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
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
	userResp, err := l.svcCtx.UserRpc.CheckEmailExist(l.ctx, &user.CheckEmailExistReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	logx.Debugf("[User] Login Password: %+v\n", userResp.User)

	if !userResp.IsExist {
		return nil, errors.New("邮箱或密码错误，请重新输入")
	}
	if userResp.User.MailStatus == 1 {
		return nil, errors.New("该邮箱不可用，请重新注册")
	}
	if userResp.User.Status != 2 {
		return nil, errors.New("该用户已被注销")
	}
	verifyCaptchaResp, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &user.VerifyCaptchaReq{
		Key:     req.CaptchaCode,
		Captcha: req.CaptchaValue,
	})
	logx.Debugf("[User] Login verifyCaptchaResp: %+v\n", verifyCaptchaResp)
	if err != nil {
		return nil, err
	}
	if !verifyCaptchaResp.IsCorrect {
		return nil, errors.New("验证码错误")
	}
	isEqual := password.VerifyPassword(userResp.User.Password, req.Password)
	if !isEqual {
		logx.Errorf("[User] Login isEqual: %v\n", isEqual)
		return nil, errors.New("邮箱或密码错误，请重新输入")
	}
	logx.Debugf("[User] Login isEqual: %+v\n", isEqual)

	secretKey := l.svcCtx.Config.JWTAuth.AccessSecret
	seconds := l.svcCtx.Config.JWTAuth.AccessExpire
	now := time.Now().Unix()
	token, err := jwtx.NewJwtToken(
		secretKey,
		seconds,
		jwtx.WithOption("userId", userResp.User.Id),
		jwtx.WithOption("nickname", userResp.User.Nickname),
		jwtx.WithOption("roleId", userResp.User.RoleId),
		jwtx.WithOption("email", userResp.User.Email),
	)
	if err != nil {
		return nil, err
	}
	if _, err = l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		TokenKey: fmt.Sprintf("%s%d", globalKey.TokenPrefixKey, userResp.User.Id),
		Token:    token,
		ExpireAt: uint64(l.svcCtx.Config.JWTAuth.AccessExpire),
	}); err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  token,
		ExpireAt:     uint64(now) + uint64(seconds),
		RefreshAfter: uint64(now) + uint64(seconds/2),
	}, nil
}
