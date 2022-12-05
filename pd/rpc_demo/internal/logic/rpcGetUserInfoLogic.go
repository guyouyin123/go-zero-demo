package logic

import (
	"context"
	"fmt"

	"go-zero-demo/pd/rpc_demo/internal/svc"
	"go-zero-demo/pd/rpc_demo/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type RpcGetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRpcGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcGetUserInfoLogic {
	return &RpcGetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RpcGetUserInfoLogic) RpcGetUserInfo(in *pb.GetUserInfoRes) (*pb.GetUserInfoReq, error) {

	if md,ok:=metadata.FromIncomingContext(l.ctx);ok{
		tmp:=md.Get("user")  //类似于网管，接收业务传过来端值
		fmt.Println("rpc客户端传过来端值：",tmp)
	}

	m := map[int64]string{
		1: "Jeff",
		2: "张三",
	}
	nickname := "unknown"
	if name, ok := m[in.Id]; ok {
		nickname = name
	}
	return &pb.GetUserInfoReq{Id: in.Id,Name: nickname}, nil
}
