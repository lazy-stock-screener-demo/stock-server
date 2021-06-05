package licenserepo

import (
	"fmt"
	customerselfdomain "stock-contexts/pkg/customer-self/domain/core"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

func TestNewReadReadRepo(t *testing.T) {
	testutils.LoadEnv()
	licenseName := customerselfdomain.NewLicenseName(
		customerselfdomain.LicenseNameProps{
			Value: "std",
		})
	fmt.Printf("Input: LicenseName %v\n", licenseName.GetValue())
	repo := NewReadRepo()
	license, err := repo.ReadLicenseByName(licenseName)
	a := license.GetLicenseID()
	fmt.Printf("Output: ID %v\n", a.GetID())
	fmt.Printf("Output: LicenseName %v\n", license.GetLicenseName().GetValue())

	if err.Error != nil {
		t.Errorf("Expected license, received %v", err)
	}
}
