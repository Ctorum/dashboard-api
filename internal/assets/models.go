package assets

import "github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/internal/infra"

type Assets struct {
	infra.TypeIDMixin
	infra.TimestampMixin
	Type               AssetType   `json:"type"`
	Status             AssetStatus `json:"status"`
	TotalPurchaseValue float32     `json:"total_purchase_value"`
}
