package createcustomer

import (
	customerdto "stock-contexts/pkg/customer-self/application/dto"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerrepo "stock-contexts/pkg/customer-self/infra/repo/gorm/Customer"
	licenserepo "stock-contexts/pkg/customer-self/infra/repo/gorm/License"

	appcore "stock-contexts/pkg/shared/application"
)

type iuseCaseCRUD interface {
	execute(req ReqDTO) (*customerdto.CustomerDTO, appcore.EitherError)
}

// UseCase interface for GetStockByTicker
type useCaseCRUD struct {
	curtomerRepo customerrepo.ICRUDCommand
	licenseRepo  licenserepo.IRead
}

func (u *useCaseCRUD) execute(req ReqDTO) (*customerdto.CustomerDTO, appcore.EitherError) {
	// Create Customer Name
	customerName := customerselfdomain.NewCustomerName(customerselfdomain.CustomerNameProps{
		Value: req.CustomerName,
	})
	if !customerName.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(customerName.Result.GetErr()),
		)
	}
	customerPWD := customerselfdomain.NewCustomerPWD(customerselfdomain.CustomerPWDProps{
		Value:  req.CustomerPWD,
		Hashed: false,
	})
	if !customerPWD.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(customerPWD.Result.GetErr()),
		)
	}
	customerEmail := customerselfdomain.NewCustomerEmail(customerselfdomain.CustomerEmailProps{
		Value: req.CustomerEmail,
	})
	if !customerEmail.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(customerEmail.Result.GetErr()),
		)
	}

	// Create License Name
	licenseName := customerselfdomain.NewLicenseName(
		customerselfdomain.LicenseNameProps{
			Value: "std",
		})
	if !licenseName.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(licenseName.Result.GetErr()),
		)
	}
	// Read License Repo
	license, err := u.licenseRepo.ReadLicenseByName(licenseName)
	if err.Error != nil {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			NewLicenseNotFoundErr(err),
		)
	}

	// Create Customer Aggregate Root
	customer := customerselfdomain.NewCustomer(nil, customerselfdomain.CustomerProps{
		CustomerName:  customerName,
		CustomerPWD:   customerPWD,
		CustomerEmail: customerEmail,
		LicenseID:     license.GetLicenseID(),
	})
	if !customer.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(customer.Result.GetErr()),
		)
	}

	// Save Customer Aggregate Root
	customerView, err := u.curtomerRepo.Save(customer)
	if err.Error != nil {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			NewCustomerSavingInRepoErr(err),
		)
	}
	// Success
	return customerView, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCaseCRUD() *useCaseCRUD {
	return &useCaseCRUD{
		curtomerRepo: customerrepo.NewCRUDCommandRepo(),
		licenseRepo:  licenserepo.NewReadRepo(),
	}
}
