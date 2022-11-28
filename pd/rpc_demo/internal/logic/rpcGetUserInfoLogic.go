package logic

import (
	"context"

	"go-zero-demo/pd/rpc_demo/internal/svc"
	"go-zero-demo/pd/rpc_demo/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

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
