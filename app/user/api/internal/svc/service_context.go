package svc

import (
	"net/http"
	"net/url"
	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/api/internal/middleware"
	"qinglv-backend/app/user/rpc/client/blacklistclass"
	"qinglv-backend/app/user/rpc/client/captchaclass"
	"qinglv-backend/app/user/rpc/client/emailclass"
	"qinglv-backend/app/user/rpc/client/followingclass"
	"qinglv-backend/app/user/rpc/client/roleclass"
	"qinglv-backend/app/user/rpc/client/userclass"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	Authority    rest.Middleware
	UserRpc      userclass.UserClass
	RoleRpc      roleclass.RoleClass
	EmailRpc     emailclass.EmailClass
	CaptchaRpc   captchaclass.CaptchaClass
	FollowingRpc followingclass.FollowingClass
	BlacklistRpc blacklistclass.BlacklistClass
	CosClient    *cos.Client
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
	svc := &ServiceContext{
		Config:       c,
		UserRpc:      userclass.NewUserClass(zrpc.MustNewClient(c.UserRpc)),
		RoleRpc:      roleclass.NewRoleClass(zrpc.MustNewClient(c.UserRpc)),
		EmailRpc:     emailclass.NewEmailClass(zrpc.MustNewClient(c.UserRpc)),
		CaptchaRpc:   captchaclass.NewCaptchaClass(zrpc.MustNewClient(c.UserRpc)),
		FollowingRpc: followingclass.NewFollowingClass(zrpc.MustNewClient(c.UserRpc)),
		BlacklistRpc: blacklistclass.NewBlacklistClass(zrpc.MustNewClient(c.UserRpc)),
		CosClient:    client,
	}
	svc.Authority = middleware.NewAuthorityMiddleware(svc.UserRpc, &c).Handle
	return svc
}
