package readautocompletebychar

type ReqDTO struct {
	StockChar string `json:"char"`
}

// ResDTO for Response DTO
type ResDTO struct {
	CreationStatus []string `json:"matchingList"`
}
