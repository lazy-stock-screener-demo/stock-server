package login

import (
	identitydto "stock-contexts/pkg/customer-identity/application/dto"
)

// ReqDTO for Request DTO
type ReqDTO struct {
	UserName     string `json:"userName"`
	UserPassword string `json:"userPWD"`
}

// ResDTO for Response DTO
type ResDTO struct {
	Identity identitydto.IdentityDTO `json:"identity"`
}
