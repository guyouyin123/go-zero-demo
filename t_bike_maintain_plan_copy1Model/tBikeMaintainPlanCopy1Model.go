package t_bike_maintain_plan_copy1Model

import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TBikeMaintainPlanCopy1Model = (*customTBikeMaintainPlanCopy1Model)(nil)

type (
	// TBikeMaintainPlanCopy1Model is an interface to be customized, add more methods here,
	// and implement the added methods in customTBikeMaintainPlanCopy1Model.
	TBikeMaintainPlanCopy1Model interface {
		tBikeMaintainPlanCopy1Model
	}

	customTBikeMaintainPlanCopy1Model struct {
		*defaultTBikeMaintainPlanCopy1Model
	}
)

// NewTBikeMaintainPlanCopy1Model returns a model for the database table.
func NewTBikeMaintainPlanCopy1Model(conn sqlx.SqlConn) TBikeMaintainPlanCopy1Model {
	return &customTBikeMaintainPlanCopy1Model{
		defaultTBikeMaintainPlanCopy1Model: newTBikeMaintainPlanCopy1Model(conn),
	}
}
