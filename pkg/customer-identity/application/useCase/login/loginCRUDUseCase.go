package login

import (
	identitydomain "stock-contexts/pkg/customer-identity/domain"
	identityrepo "stock-contexts/pkg/customer-identity/infra/repo/gorm"
	cacherepo "stock-contexts/pkg/customer-identity/infra/repo/redis"
	appcore "stock-contexts/pkg/shared/application"
	authservice "stock-contexts/pkg/shared/services/auth"
)

// IUseCaseCRUD define usecase interface
type IUseCaseCRUD interface {
	Execute(req ReqDTO) (*identitydomain.User, appcore.EitherError)
}

// UseCaseCRUD interface for GetStockByTicker
type UseCaseCRUD struct {
	authService  *authservice.JWTAuth
	cacheRepo    cacherepo.IRedisRepo
	identityRepo identityrepo.IRead
}

// Execute defines execution
func (u *UseCaseCRUD) Execute(req ReqDTO) (*identitydomain.User, appcore.EitherError) {
	// username
	userName := identitydomain.NewUserName(identitydomain.UserNameProps{
		Value: req.UserName,
	})
	if !userName.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(userName.Result.GetErr()),
		)
	}
	// password
	password := identitydomain.NewUserPWD(identitydomain.UserPWDProps{
		Value: req.UserPassword,
	})
	if !password.Result.IsSuccess() {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(password.Result.GetErr()),
		)
	}
	// get user from repo
	user, err := u.identityRepo.ReadCustomerByName(userName)
	if err.Error != nil {
		return nil, appcore.NewEitherErr(
			appcore.NewErr(),
			NewUserNotFoundErr(err),
		)
	}
	// create access token and refresh token
	accessToken, _ := u.authService.SignJWT(userName.GetValue())
	refreshToekn := u.authService.CreateRefreshToken()

	// settoken on Aggregate root
	user.SetAccessToken(accessToken, refreshToekn)

	// save jwt in redis to record login state
	u.cacheRepo.SaveAuthenticatedUser(user)

	return &user, appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

// NewUseCaseCRUD defines
func NewUseCaseCRUD() *UseCaseCRUD {
	return &UseCaseCRUD{
		authService:  authservice.NewJWTAuth(),
		cacheRepo:    cacherepo.NewCRUDRepo(),
		identityRepo: identityrepo.NewReadRepo(),
	}
}
