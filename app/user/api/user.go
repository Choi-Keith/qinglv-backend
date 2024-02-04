package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	"qinglv-backend/app/user/api/internal/config"
	"qinglv-backend/app/user/api/internal/handler"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/common/response"

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

	server := rest.MustNewServer(c.RestConf,
		rest.WithCors("*"),
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			response.FailCodeMsg(w, http.StatusUnauthorized, err)
		}), rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response.FailCodeMsg(w, http.StatusMethodNotAllowed, errors.New("请求方法出错"))
		})), rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response.FailCodeMsg(w, http.StatusNotFound, errors.New("找不到路径"))
		})))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
