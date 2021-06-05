package isauthed

import (
	// identityrepo "stock-contexts/pkg/customer-identity/infra/repo/gorm"

	identitydomain "stock-contexts/pkg/customer-identity/domain"
	cacherepo "stock-contexts/pkg/customer-identity/infra/repo/redis"
	appcore "stock-contexts/pkg/shared/application"
	authservice "stock-contexts/pkg/shared/services/auth"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// IUseCaseCRUD define usecase interface
type IUseCaseCRUD interface {
	execute(req ReqDTO) (string, appcore.EitherError)
}

type useCaseCRUD struct {
	authService *authservice.JWTAuth
	cacheRepo   cacherepo.IRedisRepo
}

func (u *useCaseCRUD) execute(req ReqDTO) (string, appcore.EitherError) {

	hasPrefix := strings.HasPrefix(req.AuthorizationHeader, "Bearer ")
	if hasPrefix == false {
		return "", appcore.NewEitherErr(
			appcore.NewErr(),
			NewNoAccessTokenProvidedErr(""),
		)
	}

	token, err := u.authService.VerifyJWT(req.AuthorizationHeader)
	if err != nil {
		return "", appcore.NewEitherErr(
			appcore.NewErr(),
			NewSignatureExpiredErr(err),
		)
	}
	// fmt.Printf("%v", token.Claims.(jwt.MapClaims)["id"].(string))
	userName := identitydomain.NewUserName(identitydomain.UserNameProps{
		Value: token.Claims.(jwt.MapClaims)["id"].(string),
	})
	if !userName.Result.IsSuccess() {
		return "", appcore.NewEitherErr(
			appcore.NewErr(),
			appcore.NewResultFailed(userName.Result.GetErr()),
		)
	}

	_, err = u.cacheRepo.ReadTokens(userName)
	if err != nil {
		return "", appcore.NewEitherErr(
			appcore.NewErr(),
			NewUserNotFoundErr(err),
		)
	}
	return userName.GetValue(), appcore.NewEitherErr(
		appcore.NewSuccess(),
		appcore.NewResultOk(),
	)
}

func NewUseCaseCRUD() *useCaseCRUD {
	return &useCaseCRUD{
		authService: authservice.NewJWTAuth(),
		cacheRepo:   cacherepo.NewCRUDRepo(),
	}
}
