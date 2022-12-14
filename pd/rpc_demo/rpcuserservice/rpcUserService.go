// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package rpcuserservice

import (
	"context"

	"go-zero-demo/pd/rpc_demo/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq = pb.GetUserInfoReq
	GetUserInfoRes = pb.GetUserInfoRes

	RpcUserService interface {
		RpcGetUserInfo(ctx context.Context, in *GetUserInfoRes, opts ...grpc.CallOption) (*GetUserInfoReq, error)
	}

	defaultRpcUserService struct {
		cli zrpc.Client
	}
)

func NewRpcUserService(cli zrpc.Client) RpcUserService {
	return &defaultRpcUserService{
		cli: cli,
	}
}

func (m *defaultRpcUserService) RpcGetUserInfo(ctx context.Context, in *GetUserInfoRes, opts ...grpc.CallOption) (*GetUserInfoReq, error) {
	client := pb.NewRpcUserServiceClient(m.cli.Conn())
	return client.RpcGetUserInfo(ctx, in, opts...)
}
