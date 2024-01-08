package user

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginCaptchaLogic {
	return &LoginCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCaptchaLogic) LoginCaptcha() (resp *types.LoginCaptchaResp, err error) {
	// todo: add your logic here and delete this line
	generateCaptchaResp, err := l.svcCtx.UserRpc.GenerateCaptcha(l.ctx, &user_client.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.LoginCaptchaResp{
		CaptchaCode: generateCaptchaResp.Key,
		CaptchaImg:  generateCaptchaResp.Base64Img,
	}, nil
}
