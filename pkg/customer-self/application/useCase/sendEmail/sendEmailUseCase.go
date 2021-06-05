package sendEmail

import (
	appcore "stock-contexts/pkg/shared/application"
)

type iuseCase interface {
	execute(req ReqDTO) ([]string, appcore.EitherError)
}

type useCase struct {
}

func (u *useCase) execute(req ReqDTO) ([]string, appcore.EitherError) {
	return []string{""}, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func newUseCase() *useCase {
	return &useCase{}
}
