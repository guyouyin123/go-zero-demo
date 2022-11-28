package bike

import (
	"go-zero-demo/api/rpc_demo/model/bikeModel"
	"context"
	"errors"

	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BikeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBikeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BikeInfoLogic {
	return &BikeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BikeInfoLogic) BikeInfo(req *types.BikeInfoRes) (resp *types.BikeInfoReq, err error) {
	// todo: add your logic here and delete this line
	//m:=map[int]string{
	//	1:"摩托车",
	//	2:"自行车",
	//}
	//bikeName:="unknown"
	//if name,ok:=m[req.BikeId];ok {
	//	bikeName = name
	//}
	//return &types.BikeInfoReq{BikeId: req.BikeId,BikeName: bikeName},nil

	bike, err := l.svcCtx.BikeDbModel.FindOne(l.ctx, int64(req.BikeId))
	if err != nil && err != bikeModel.ErrNotFound {
		return nil, errors.New("查询数据库报错！")
	}
	if bike == nil {
		return nil, errors.New("车辆不存在！")
	}
	return &types.BikeInfoReq{BikeId: req.BikeId,BikeName: bike.BikeName},nil
}
