package createcustomer

import (
	customerselfdto "stock-contexts/pkg/customer-self/application/dto"
)

// ReqDTO for Request DTO
type ReqDTO struct {
	CustomerName  string `json:"userName"`
	CustomerPWD   string `json:"userPWD"`
	CustomerEmail string `json:"email"`
}

// ResDTO for Response DTO
type ResDTO struct {
	Customer *customerselfdto.CustomerDTO `json:"customer"`
}
