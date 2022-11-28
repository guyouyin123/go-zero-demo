package user

import (
	"go-zero-demo/api/rpc_demo/model/userModel"
	"context"

	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRes) (resp *types.UserInfoReq, err error) {
	// todo: add your logic here and delete this line

	//m := map[int]string{
	//	1: "Jeff",
	//	2: "张三",
	//}
	//nickname := "unknown"
	//if name, ok := m[req.UserId]; ok {
	//	nickname = name
	//}
	//return &types.UserInfoReq{UserId: req.UserId, Name: user.nickname}, nil
	logx.Info("===========user log info=========")
	errors.New("故意报错测试堆栈")

	user, err := l.svcCtx.UserDbModel.FindOne(l.ctx, int64(req.UserId))
	if err != nil && err!= userModel.ErrNotFound {
		return nil, errors.New("查询数据库报错！")
	}
	if user ==nil{
		return nil, errors.New("用户不存在！")
	}

	return &types.UserInfoReq{UserId: req.UserId, Name: user.Name}, nil
}
