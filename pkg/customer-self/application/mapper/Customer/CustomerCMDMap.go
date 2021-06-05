package customermapper

import (
	customerdto "stock-contexts/pkg/customer-self/application/dto"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerselfschema "stock-contexts/pkg/customer-self/infra/repo/schema"
	idcore "stock-contexts/pkg/shared/domain/id"
	hashhelper "stock-contexts/pkg/shared/utils/ecrypt"
)

// EntityMap struct
type EntityMap struct{}

// ToPersistence method
func (e EntityMap) ToPersistence(cus customerselfdomain.Customer) *customerselfschema.Customer {
	customerID := cus.GetCustomerID()
	licenseID := cus.GetLicenseID()
	cusUUID, _ := customerID.GetUUID()
	licenseUUID, _ := licenseID.GetUUID()
	hashedPWD, _ := hashhelper.GetHashedValue(cus.GetCustomerPWD().GetValue())
	return &customerselfschema.Customer{
		ID:           cusUUID,
		CustomerName: cus.GetCustomerName().GetValue(),
		CustomerPWD:  hashedPWD,
		Email:        cus.GetEmail().GetValue(),
		LicenseID:    licenseUUID,
	}
}

// ToDomain Implemented
func (e EntityMap) ToDomain(raw *customerselfschema.Customer) customerselfdomain.Customer {
	customerName := customerselfdomain.NewCustomerName(
		customerselfdomain.CustomerNameProps{
			Value: raw.CustomerName,
		})
	licenseID := customerselfdomain.NewLicenseID(
		&idcore.UniqueEntityID{
			Identifier: &idcore.Identifier{
				Value: raw.LicenseID.String(),
			},
		})
	customer := customerselfdomain.NewCustomer(
		idcore.NewIDEntity(raw.ID.String()),
		customerselfdomain.CustomerProps{
			CustomerName: customerName,
			LicenseID:    licenseID,
		})
	return customer
}

// ToDTO method
func (e EntityMap) ToDTO(customerView *customerselfschema.Customer) customerdto.CustomerDTO {
	return customerdto.CustomerDTO{
		CustomerName:  customerView.CustomerName,
		CustomerEmail: customerView.Email,
	}
}

// NewEntityMap method
func NewEntityMap() EntityMap {
	return EntityMap{}
}
