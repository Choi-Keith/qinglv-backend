package main

import (
	"flag"
	"fmt"

	"qinglv-backend/app/operation/rpc/internal/config"
	collectionclassServer "qinglv-backend/app/operation/rpc/internal/server/collectionclass"
	commentclassServer "qinglv-backend/app/operation/rpc/internal/server/commentclass"
	shareclassServer "qinglv-backend/app/operation/rpc/internal/server/shareclass"
	thumbclassServer "qinglv-backend/app/operation/rpc/internal/server/thumbclass"
	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/operation.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		operation.RegisterCollectionClassServer(grpcServer, collectionclassServer.NewCollectionClassServer(ctx))
		operation.RegisterShareClassServer(grpcServer, shareclassServer.NewShareClassServer(ctx))
		operation.RegisterThumbClassServer(grpcServer, thumbclassServer.NewThumbClassServer(ctx))
		operation.RegisterCommentClassServer(grpcServer, commentclassServer.NewCommentClassServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
