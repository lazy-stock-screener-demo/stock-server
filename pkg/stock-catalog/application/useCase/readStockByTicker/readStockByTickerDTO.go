package readstockbyticker

import (
	catalogdto "stock-contexts/pkg/stock-catalog/application/dto"
)

// ReqDTO for Request DTO
type ReqDTO struct {
	StockVID string `json:"stockVID"`
}

// ResDTO for Response DTO
type ResDTO struct {
	Catalog catalogdto.CatalogDTO `json:"catalog"`
}
