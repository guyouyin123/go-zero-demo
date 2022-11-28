package user

import (
	"context"
	"go-zero-demo/pd/rpc_demo/pb"

	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RpcUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRpcUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcUserInfoLogic {
	return &RpcUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//api的RPC调用
func (l *RpcUserInfoLogic) RpcUserInfo(req *types.UserInfoRes) (resp *types.UserInfoReq, err error) {
	// todo: add your logic here and delete this line

	userRes,err:= l.svcCtx.RpcUserRpcClient.RpcGetUserInfo(l.ctx,&pb.GetUserInfoRes{Id: int64(req.UserId)})
	if err !=nil{
		return nil, err
	}
	return &types.UserInfoReq{
		UserId: int(userRes.Id),
		Name:   userRes.Name,
	},nil
}
