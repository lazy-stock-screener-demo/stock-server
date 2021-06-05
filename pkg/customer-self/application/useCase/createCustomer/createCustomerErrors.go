package createcustomer

import (
	appcore "stock-contexts/pkg/shared/application"
)

// NewLicenseNotFoundErr define implicit error type
func NewLicenseNotFoundErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't read a license with license name",
			Type:    "LicenseNotFound",
			Error:   err,
		},
	)
}

// NewCustomerSavingInRepoErr define implicit error type
func NewCustomerSavingInRepoErr(err interface{}) appcore.Result {
	return appcore.NewResultUseCaseErr(
		appcore.ErrProps{
			Message: "Couldn't create a customer with customer name",
			Type:    "CustomerSavingInRepo",
			Error:   err,
		},
	)
}
