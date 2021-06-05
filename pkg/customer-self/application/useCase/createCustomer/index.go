package createcustomer

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var conCRUD = newControllerCRUD()

// var conCQRS = newControllerCQRS()
// var gqlcon = newGraphQLController()

// ExecuteCRUD return a function to fasthttp router
func ExecuteCRUD() func(c *routing.Context) error {
	return func(c *routing.Context) error {
		conCRUD.executeImpl(c)
		return nil
	}
}

// // ExecuteCQRS return a function to fasthttp router
// func ExecuteCQRS() func(c *routing.Context) error {
// 	return func(c *routing.Context) error {
// 		conCQRS.executeImpl(c)
// 		return nil
// 	}
// }

// // GraphQLExec return a graphql to fasthttp router
// func GraphQLExec(parms graphql.ResolveParams) (interface{}, error) {
// 	result, _ := gqlcon.executeImpl(parms)
// 	return result, nil
// }
