package svc

import (
	"go-zero-demo/api/rpc_demo/internal/config"
	"go-zero-demo/api/rpc_demo/internal/middleware"
	"go-zero-demo/api/rpc_demo/model/bikeModel"
	"go-zero-demo/api/rpc_demo/model/userModel"
	"go-zero-demo/pd/rpc_demo/rpcuserservice"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware //局部中间件
	UserDbModel    userModel.UsersModel
	BikeDbModel    bikeModel.BikeModel
	RpcUserRpcClient  rpcuserservice.RpcUserService //Rpc配置
}

//初始化
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle, //局部中间件
		UserDbModel:    userModel.NewUsersModel(sqlx.NewMysql(c.MysqlDB.DataSource)),
		BikeDbModel:    bikeModel.NewBikeModel(sqlx.NewMysql(c.MysqlDB.DataSource), c.Redis),
		RpcUserRpcClient:  rpcuserservice.NewRpcUserService(zrpc.MustNewClient(c.UserRpcConf)), //Rpc配置初始化
	}
}
