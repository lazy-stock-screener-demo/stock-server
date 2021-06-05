package readstockbyticker

import (
	appcore "stock-contexts/pkg/shared/application"
	catalogdomain "stock-contexts/pkg/stock-catalog/domain"
	catalogrepo "stock-contexts/pkg/stock-catalog/infra/repo/mongodriver"
)

type iuseCase interface {
	execute(req ReqDTO) (catalogdomain.StockView, appcore.EitherError)
}

// UseCase interface for GetStockByTicker
type useCase struct {
	catalogrepo catalogrepo.IRead
}

func (u *useCase) execute(req ReqDTO) (s catalogdomain.StockView, e appcore.EitherError) {
	stockView, err := u.catalogrepo.ReadStocksByTicker(req.StockVID)
	if err.Error != nil {
		return stockView, appcore.NewEitherErr(
			appcore.NewErr(),
			NewTickerNotFoundErr(err),
		)
	}
	return stockView, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCase() *useCase {
	return &useCase{
		catalogrepo: catalogrepo.NewReadRepo(),
	}
}

// defer func() {
// 	if err := recover(); err != nil {
// 		fmt.Println(err)
// 	}
// }()
