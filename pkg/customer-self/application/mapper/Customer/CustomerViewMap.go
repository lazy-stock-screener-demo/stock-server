package customermapper

import (
	customerdto "stock-contexts/pkg/customer-self/application/dto"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
)

type ViewMap struct{}

// ToDomain Implemented
func (v ViewMap) ToDomain(raw *customerselfschema.Customer) *customerselfdomain.CustomerView {
	customerName := customerselfdomain.NewCustomerName(
		customerselfdomain.CustomerNameProps{
			Value: raw.CustomerName,
		},
	)

	customerView := customerselfdomain.NewCustomerView(
		customerselfdomain.CustomerViewProps{
			CustomerName: customerName,
		},
	)
	return customerView
}

// ToDTO method
func (v ViewMap) ToDTO(customerView *customerselfschema.CustomerView) customerdto.CustomerDTO {
	return customerdto.CustomerDTO{
		CustomerName:        customerView.CustomerName,
		CustomerEmail:       customerView.Email,
		CustomerLicenseName: customerView.LicenseName,
	}
}

// NewViewMap define
func NewViewMap() ViewMap {
	return ViewMap{}
}
