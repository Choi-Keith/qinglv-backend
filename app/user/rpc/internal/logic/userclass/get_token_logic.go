package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenLogic {
	return &GetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenLogic) GetToken(in *user.GetTokenReq) (*user.GetTokenResp, error) {
	// todo: add your logic here and delete this line
	token, err := l.svcCtx.RedisClient.Get(in.TokenKey)
	if err != nil {
		return &user.GetTokenResp{
			Token: "",
		}, err
	}
	return &user.GetTokenResp{
		Token: token,
	}, nil
}
