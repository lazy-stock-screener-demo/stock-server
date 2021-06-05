package createcustomer

import (
	"fmt"
	customerrepo "stock-contexts/pkg/customer-self/infra/repo/gorm/Customer"
	licenserepo "stock-contexts/pkg/customer-self/infra/repo/gorm/License"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

func TestNewUseCaseCRUD(t *testing.T) {
	testutils.LoadEnv()
	useCase := &useCaseCRUD{
		curtomerRepo: customerrepo.NewCRUDCommandRepo(),
		licenseRepo:  licenserepo.NewReadRepo(),
	}
	customerView, eitherErr := useCase.execute(ReqDTO{
		CustomerName: "test",
	})
	fmt.Printf("Output: Customer Name %v\n", customerView)

	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		t.Errorf("Error: Expected no Error, received error: %v", err)
	}
}
