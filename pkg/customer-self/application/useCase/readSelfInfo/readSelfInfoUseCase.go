package readselfinfo

import (
	customerdto "stock-contexts/pkg/customer-self/application/dto"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	customerrepo "stock-contexts/pkg/customer-self/infra/repo/gorm/Customer"
	appcore "stock-contexts/pkg/shared/application"
)

type iuseCase interface {
	execute(req ReqDTO) (*customerdto.CustomerDTO, appcore.EitherError)
}

// UseCase interface for GetStockByTicker
type useCase struct {
	customerrepo customerrepo.IRead
}

func (u *useCase) execute(req ReqDTO) (*customerdto.CustomerDTO, appcore.EitherError) {
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
	customerView, err := u.customerrepo.ReadCustomerViewByName(customerName)
	if err.Error != nil {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			NewCustomerNotFoundErr(err),
		)
	}
	return customerView, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCase() *useCase {
	return &useCase{
		customerrepo: customerrepo.NewReadRepo(),
	}
}
