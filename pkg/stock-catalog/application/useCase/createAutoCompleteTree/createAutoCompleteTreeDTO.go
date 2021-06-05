package autocompletesearch

type ReqDTO struct {
	StockVID string `json:"stockVID"`
}

// ResDTO for Response DTO
type ResDTO struct {
	CreationStatus string `json:"creationStatus"`
}
