package readstockbyticker

import (
	"github.com/graphql-go/graphql"
	routing "github.com/qiangxue/fasthttp-routing"
)

var con = newController()
var gqlcon = newGraphQLController()

// Execute return a function to fasthttp router
func Execute() func(c *routing.Context) error {
	return func(c *routing.Context) error {
		con.executeImpl(c)
		return nil
	}
}

// GraphQLExec return a graphql to fasthttp router
func GraphQLExec(parms graphql.ResolveParams) (interface{}, error) {
	result, _ := gqlcon.executeImpl(parms)
	return result, nil
}
