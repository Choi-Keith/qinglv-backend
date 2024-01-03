package main

import (
	"flag"
	"fmt"

	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/api/internal/handler"
	"qinglv-backend/app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	logConf := logx.LogConf{
		Mode:  "console",
		Stat:  false,
		Level: "debug",
	}
	logx.MustSetup(logConf)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("127.0.0.1"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
