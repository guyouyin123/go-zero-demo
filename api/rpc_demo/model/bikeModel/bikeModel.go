package bikeModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BikeModel = (*customBikeModel)(nil)

type (
	// BikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBikeModel.
	BikeModel interface {
		bikeModel
	}

	customBikeModel struct {
		*defaultBikeModel
	}
)

// NewBikeModel returns a model for the database table.
func NewBikeModel(conn sqlx.SqlConn, c cache.CacheConf) BikeModel {
	return &customBikeModel{
		defaultBikeModel: newBikeModel(conn, c),
	}
}
