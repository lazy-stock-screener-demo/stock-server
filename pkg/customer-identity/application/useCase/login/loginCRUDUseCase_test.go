package login

import (
	"fmt"
	identityrepo "stock-contexts/pkg/customer-identity/infra/repo/gorm"
	cacherepo "stock-contexts/pkg/customer-identity/infra/repo/redis"
	authservice "stock-contexts/pkg/shared/services/auth"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

func TestLoginCRUDUseCase(t *testing.T) {
	testutils.LoadEnv()
	useCase := &UseCaseCRUD{
		authService:  authservice.NewJWTAuth(),
		cacheRepo:    cacherepo.NewCRUDRepo(),
		identityRepo: identityrepo.NewReadRepo(),
	}
	user, eitherErr := useCase.Execute(ReqDTO{
		UserName:     "test1",
		UserPassword: "love123",
	})

	fmt.Printf("Output: user Login %v\n", user.GetUserName().GetValue())
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		t.Errorf("Error: Expected no Error, received error: %v", err)
	}
}
