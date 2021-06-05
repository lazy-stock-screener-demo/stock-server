package catalogdto

// SafetyDTO define DTO of Profit part
type SafetyDTO struct {
	LongTermDebt []map[string]float32 `json:"longTermDebt"`
	CurrentRatio []map[string]float64 `json:"currentRatio"`
}
