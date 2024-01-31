package main

import (
	"flag"
	"fmt"

	"qinglv-backend/app/user/rpc/internal/config"
	blacklistclassServer "qinglv-backend/app/user/rpc/internal/server/blacklistclass"
	captchaclassServer "qinglv-backend/app/user/rpc/internal/server/captchaclass"
	emailclassServer "qinglv-backend/app/user/rpc/internal/server/emailclass"
	followingclassServer "qinglv-backend/app/user/rpc/internal/server/followingclass"
	roleclassServer "qinglv-backend/app/user/rpc/internal/server/roleclass"
	userclassServer "qinglv-backend/app/user/rpc/internal/server/userclass"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	logConf := logx.LogConf{
		Mode:  "console",
		Stat:  false,
		Level: "debug",
	}
	logx.MustSetup(logConf)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterEmailClassServer(grpcServer, emailclassServer.NewEmailClassServer(ctx))
		user.RegisterCaptchaClassServer(grpcServer, captchaclassServer.NewCaptchaClassServer(ctx))
		user.RegisterRoleClassServer(grpcServer, roleclassServer.NewRoleClassServer(ctx))
		user.RegisterUserClassServer(grpcServer, userclassServer.NewUserClassServer(ctx))
		user.RegisterFollowingClassServer(grpcServer, followingclassServer.NewFollowingClassServer(ctx))
		user.RegisterBlacklistClassServer(grpcServer, blacklistclassServer.NewBlacklistClassServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
