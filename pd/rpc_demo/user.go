package main

import (
	"flag"
	"fmt"

	"go-zero-demo/pd/rpc_demo/internal/config"
	"go-zero-demo/pd/rpc_demo/internal/server"
	"go-zero-demo/pd/rpc_demo/internal/svc"
	"go-zero-demo/pd/rpc_demo/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "/Users/jeff/myself/go-zero-demo/pd/rpc_demo/etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRpcUserServiceServer(grpcServer, server.NewRpcUserServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
