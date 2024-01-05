package config

import (
	"qinglv-backend/pkg/email"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type WebsiteOptions struct {
	Host string
	Port int
}

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		Datasource string
	}

	Cache cache.CacheConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Website WebsiteOptions

	SMTP email.SMTPOptions
}
