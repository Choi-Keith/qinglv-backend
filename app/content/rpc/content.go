package main

import (
	"flag"
	"fmt"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/config"
	categoryclassServer "qinglv-backend/app/content/rpc/internal/server/categoryclass"
	mediafileclassServer "qinglv-backend/app/content/rpc/internal/server/mediafileclass"
	postclassServer "qinglv-backend/app/content/rpc/internal/server/postclass"
	topicclassServer "qinglv-backend/app/content/rpc/internal/server/topicclass"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/content.yaml", "the config file")

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
		content.RegisterTopicClassServer(grpcServer, topicclassServer.NewTopicClassServer(ctx))
		content.RegisterCategoryClassServer(grpcServer, categoryclassServer.NewCategoryClassServer(ctx))
		content.RegisterPostClassServer(grpcServer, postclassServer.NewPostClassServer(ctx))
		content.RegisterMediaFileClassServer(grpcServer, mediafileclassServer.NewMediaFileClassServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
