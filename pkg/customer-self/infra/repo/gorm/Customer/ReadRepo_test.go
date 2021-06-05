package customerrepo

import (
	"fmt"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	hashhelper "stock-contexts/pkg/shared/utils/ecrypt"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

func TestNewReadReadRepo(t *testing.T) {
	testutils.LoadEnv()
	customerName := customerselfdomain.NewCustomerName(
		customerselfdomain.CustomerNameProps{
			Value: "test1",
		})
	fmt.Printf("Input: CustomerName %v\n", customerName.GetValue())
	repo := NewReadRepo()
	customerView, err := repo.ReadCustomerViewByName(customerName)
	fmt.Printf("Output: CustomerView %v\n", customerView)

	if err.Error != nil {
		t.Errorf("Expected customer, received %v", err)
	}
}

func TestPassword(t *testing.T) {
	hashed, err := hashhelper.GetHashedValue("love123")
	password := customerselfdomain.NewCustomerPWD(customerselfdomain.CustomerPWDProps{
		Value:  string(hashed),
		Hashed: true,
	})
	fmt.Printf("Password Bcrypt test: %v \n", password.ComparePassword("love123"))
	if err != nil {
		t.Errorf("Password Test Failed")
	}
}
