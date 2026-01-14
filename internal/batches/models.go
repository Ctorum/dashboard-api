package batches

import (
	"time"

	"github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"
)

type Batches struct {
	infra.TypeIDMixin
	infra.TimestampMixin
	FundId         string    `json:"fund_id"`
	AssignmentDate time.Time `json:"assignment_date"`
	Status         Status    `json:"status"`
}

func (b Batches) TableName() string {
	return "batch.batches"
}
