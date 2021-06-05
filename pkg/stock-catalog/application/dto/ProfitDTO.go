package catalogdto

// ProfitDTO define DTO of Profit part
type ProfitDTO struct {
	NetIncome         []map[string]float32 `json:"netIncome"`
	OperatingCashFlow []map[string]float32 `json:"operatingCashFlow"`
}
