package funds

import "github.com/riza-asset/ram-ng-sec-fidc-dashboard-api/api/infra"

type Funds struct {
	infra.UUIDMixin
	infra.TimestampMixin
	Alias                   string     `json:"alias"`
	Name                    string     `json:"name"`
	Cnpj                    string     `json:"cnpj"`
	AssetType               Asset      `gorm:"type:assettypeconcept" json:"asset_type"`
	Category                Category   `gorm:"type:categoryconcept" json:"category"`
	Originator              Originator `gorm:"type:originatorconcept" json:"originator"`
	Debtor                  Debtor     `gorm:"type:debtorconcept" json:"debtor"`
	AssetPricing            bool       `json:"asset_pricing"`
	ReceivablesRegistration bool       `json:"receivables_registration"`
}

func (f Funds) TableName() string {
	return "fund.funds"
}
