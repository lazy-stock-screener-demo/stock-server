package catalogrepo

import (
	"fmt"
	testutils "stock-contexts/pkg/shared/utils/test"
	"testing"
)

func TestReadStocksByTicker(t *testing.T) {
	testutils.LoadEnv()
	r := NewReadRepo()
	total := 1
	domain, _ := r.ReadStocksByTicker("MSFT")
	fmt.Print(domain)
	// print(err)
	if total != 1 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
