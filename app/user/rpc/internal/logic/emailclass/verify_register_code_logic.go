package emailclasslogic

import (
	"context"
	"fmt"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyRegisterCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyRegisterCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyRegisterCodeLogic {
	return &VerifyRegisterCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: email
func (l *VerifyRegisterCodeLogic) VerifyRegisterCode(in *user.VerifyRegisterCodeReq) (*user.VerifyRegisterCodeResp, error) {
	// todo: add your logic here and delete this line
	key := fmt.Sprintf("%s%s", globalKey.VerifyEmailCodePrefixKey, in.Code)

	codeContentStr, err := l.svcCtx.RedisClient.Get(key)
	if err != nil {
		return nil, err
	}

	return &user.VerifyRegisterCodeResp{
		CodeContent: codeContentStr,
	}, nil
}
