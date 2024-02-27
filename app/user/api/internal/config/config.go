package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc    zrpc.RpcClientConf
	ContentRpc zrpc.RpcClientConf

	JWTAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Cos struct {
		Endpoint      string
		Service       string
		SecretID      string
		SecretKey     string
		AvatarPath    string
		ProfileBgPath string
	}
}
