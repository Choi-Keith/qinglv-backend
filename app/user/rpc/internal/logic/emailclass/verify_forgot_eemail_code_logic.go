package emailclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyForgotEemailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyForgotEemailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyForgotEemailCodeLogic {
	return &VerifyForgotEemailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyForgotEemailCodeLogic) VerifyForgotEemailCode(in *user.VerifyForgotEemailCodeReq) (*user.VerifyRegisterCodeResp, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf("%s%s", globalKey.VerifyForgotPasswordEmailCodePrefixKey, in.Code)

	codeContentStr, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		return nil, err
	}
	l.svcCtx.RedisClient.Del(key)

	return &user.VerifyRegisterCodeResp{
		CodeContent: codeContentStr,
	}, nil
}
