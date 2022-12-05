package svc

import (
	"context"
	"fmt"
	"go-zero-demo/api/rpc_demo/internal/config"
	"go-zero-demo/api/rpc_demo/internal/middleware"
	"go-zero-demo/api/rpc_demo/model/bikeModel"
	"go-zero-demo/api/rpc_demo/model/userModel"
	"go-zero-demo/pd/rpc_demo/rpcuserservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceContext struct {
	Config           config.Config
	TestMiddleware   rest.Middleware //局部中间件
	UserDbModel      userModel.UsersModel
	BikeDbModel      bikeModel.BikeModel
	RpcUserRpcClient rpcuserservice.RpcUserService //Rpc配置
}

//初始化
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle, //局部中间件
		UserDbModel:    userModel.NewUsersModel(sqlx.NewMysql(c.MysqlDB.DataSource)),
		BikeDbModel:    bikeModel.NewBikeModel(sqlx.NewMysql(c.MysqlDB.DataSource), c.Redis),
		//RpcUserRpcClient:  rpcuserservice.NewRpcUserService(zrpc.MustNewClient(c.UserRpcConf)), //Rpc配置初始化
		RpcUserRpcClient: rpcuserservice.NewRpcUserService(zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(TestInterceptor))), //Rpc配置初始化，并添加rpc客户端拦截器
	}
}

func TestInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("====rpc客户端拦截器之前====")
	fmt.Println("====req:", req)
	fmt.Println("====method:", method)

	md := metadata.New(map[string]string{"user": "Jeff"}) //传值,类似于网管传业务需要端数据
	ctx = metadata.NewOutgoingContext(ctx, md)

	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	fmt.Println("====rpc客户端拦截器之后====")
	return nil
}
