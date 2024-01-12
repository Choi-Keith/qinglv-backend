package svc

import (
	"net/http"
	"net/url"
	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   user_client.User
	CosClient *cos.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	u, _ := url.Parse(c.Cos.Endpoint)
	su, _ := url.Parse(c.Cos.Service)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  c.Cos.SecretID,
			SecretKey: c.Cos.SecretKey,
		},
	})
	return &ServiceContext{
		Config:    c,
		UserRpc:   user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CosClient: client,
	}
}
