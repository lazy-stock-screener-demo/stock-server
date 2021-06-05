package autocompletesearch

import (
	appcore "stock-contexts/pkg/shared/application"
	ternarySearchTreeService "stock-contexts/pkg/shared/services/ternarySearchTree"
	catalogrepo "stock-contexts/pkg/stock-catalog/infra/repo/mongodriver"
)

type iuseCase interface {
	execute() (*ternarySearchTreeService.Tree, appcore.EitherError)
}

// UseCase interface for GetStockByTicker
type useCase struct {
	catalogrepo catalogrepo.IRead
	TSTService  *ternarySearchTreeService.Tree
}

func (u *useCase) execute() (*ternarySearchTreeService.Tree, appcore.EitherError) {
	tickerList, err := u.catalogrepo.ReadStockTickerList()
	// Build tree
	u.TSTService.Init(tickerList)
	if err.Error != nil {
		return u.TSTService, appcore.NewEitherErr(
			appcore.NewErr(),
			NewTickerNotFoundErr(err),
		)
	}
	return u.TSTService, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCase() *useCase {
	return &useCase{
		catalogrepo: catalogrepo.NewReadRepo(),
		TSTService:  ternarySearchTreeService.NewTree(),
	}
}
