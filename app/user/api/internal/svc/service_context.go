package svc

import (
	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
