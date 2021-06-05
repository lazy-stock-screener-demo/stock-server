package readselfinfo

import (
	customerselfdto "stock-contexts/pkg/customer-self/application/dto"
)

// ReqDTO for Request DTO
type ReqDTO struct {
	CustomerName string `json:"name"`
}

// ResDTO for Response DTO
type ResDTO struct {
	Customer *customerselfdto.CustomerDTO `json:"customer"`
}
