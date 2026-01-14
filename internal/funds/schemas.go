package funds

type BasicFundSchema struct {
	Alias     string `json:"alias" gorm:"unique"`
	Name      string `json:"name"`
	Cnpj      string `json:"cnpj"`
	AssetType string `json:"asset_type"`
}
