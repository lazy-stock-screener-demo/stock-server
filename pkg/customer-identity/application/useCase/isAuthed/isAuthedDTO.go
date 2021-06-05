package isauthed

// ReqDTO for Request DTO
type ReqDTO struct {
	// UserName            string `json:"userName"`
	AuthorizationHeader string `json:"authorizationHeader"`
}

type ResDTO struct {
	UserName string `json:"userName"`
}
