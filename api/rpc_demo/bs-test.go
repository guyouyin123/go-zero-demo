package main

import (
	"flag"
	"fmt"
	"go-zero-demo/api/rpc_demo/internal/common/middleware"

	"go-zero-demo/api/rpc_demo/internal/config"
	"go-zero-demo/api/rpc_demo/internal/handler"
	"go-zero-demo/api/rpc_demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "/Users/jeff/myself/go-zero-demo/api/rpc_demo/etc/bs-Test.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewGlobalMiddleware().Handle) //全局中间件
	logx.DisableStat() //禁用每分钟打印的监控日志，建议打开

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
