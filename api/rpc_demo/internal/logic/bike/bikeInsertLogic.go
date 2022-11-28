package bike

import (
	"go-zero-demo/api/rpc_demo/model/bikeModel"
	"context"
	"database/sql"
	"fmt"

	"go-zero-demo/api/rpc_demo/internal/svc"
	"go-zero-demo/api/rpc_demo/internal/types"

	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type BikeInsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBikeInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BikeInsertLogic {
	return &BikeInsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BikeInsertLogic) BikeInsert(req *types.BikeInfoReq) (resp *types.Res, err error) {
	// todo: add your logic here and delete this line
	data:=bikeModel.Bike{
		//BikeId:   int64(req.BikeId),
		BikeId:   sql.NullInt64{int64(req.BikeId),true},
		BikeName: req.BikeName,
	}

	Result, err:=l.svcCtx.BikeDbModel.Insert(l.ctx,&data)
	if err!=nil{
		return nil, errors.New("添加失败！")
	}

	fmt.Println(Result)
	return &types.Res{
		State:     "00",
		StateInfo: "ok!",
	}, nil
}
