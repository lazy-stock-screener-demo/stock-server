package catalogdto

// CatalogDTO struct
type CatalogDTO struct {
	StockTicker   string    `json:"ticker"`
	ProfitDetails ProfitDTO `json:"profitDetails"`
	SafetyDetails SafetyDTO `json:"safetyDetails"`
}
