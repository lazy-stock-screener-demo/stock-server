package customerrestful

import (
	isAuthed "stock-contexts/pkg/customer-identity/application/useCase/isAuthed"
	createCustomer "stock-contexts/pkg/customer-self/application/useCase/createCustomer"
	readCustomerSelf "stock-contexts/pkg/customer-self/application/useCase/readSelfInfo"

	routing "github.com/qiangxue/fasthttp-routing"
)

// var (
// 	corsAllowHeaders     = "authorization"
// 	corsAllowMethods     = "HEAD, GET,POST,PUT,DELETE,OPTIONS"
// 	corsAllowOrigin      = "*"
// 	corsAllowCredentials = "true"
// )

// ctx.Request.WriteTo(os.Stdout)

// MuxRouter for stock-catalog
func MuxRouter(api *routing.RouteGroup) {
	// api.Get("/customer", middleware.Authenticating(), readCustomerSelf.ExecuteCRUD())
	api.Get("/customer", isAuthed.ExecuteCRUD(), readCustomerSelf.ExecuteCRUD())
	api.Post("/signup", createCustomer.ExecuteCRUD())
}
