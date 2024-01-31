package captchaclasslogic

import (
	"context"
	"strings"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: captcha
func (l *VerifyCaptchaLogic) VerifyCaptcha(in *user.VerifyCaptchaReq) (*user.VerifyCaptchaResp, error) {
	// todo: add your logic here and delete this line
	realCaptcha, err := l.svcCtx.RedisClient.Get(in.Key)
	if err != nil {
		return nil, err
	}
	if _, err := l.svcCtx.RedisClient.Del(in.Key); err != nil {
		return nil, err
	}
	flag := false
	if strings.TrimSpace(in.Captcha) == realCaptcha {
		flag = true
	}
	return &user.VerifyCaptchaResp{
		IsCorrect: flag,
	}, nil
}
