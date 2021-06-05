package customerdto

type CustomerDTO struct {
	CustomerName        string `json:"customerName"`
	CustomerEmail       string `json:"email"`
	CustomerLicenseName string `json:"license"`
}
