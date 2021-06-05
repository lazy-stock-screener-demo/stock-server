package login

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var conCRUD = newControllerCRUD()

// var conCQRS = newControllerCQRS()
// var gqlcon = newGraphQLController()

// ExecuteCRUD return a function to fasthttp router
func ExecuteCRUD() func(*routing.Context) error {
	return func(c *routing.Context) error {
		conCRUD.executeImpl(c)
		return nil
	}
}
