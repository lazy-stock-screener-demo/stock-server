package stockcataloggraphql

import (
	readstockbyticker "stock-contexts/pkg/stock-catalog/application/useCase/readStockByTicker"

	"github.com/graphql-go/graphql"
)

func readStockByTicker(p graphql.ResolveParams) (interface{}, error) {
	result, _ := readstockbyticker.GraphQLExec(p)
	return result, nil
}
