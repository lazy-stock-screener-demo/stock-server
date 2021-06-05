package identityrestful

import (
	isAuthed "stock-contexts/pkg/customer-identity/application/useCase/isAuthed"
	login "stock-contexts/pkg/customer-identity/application/useCase/login"

	routing "github.com/qiangxue/fasthttp-routing"
)

// middleware.Authorizing()

// var (
// 	corsAllowHeaders     = "authorization"
// 	corsAllowMethods     = "HEAD, GET,POST,PUT,DELETE,OPTIONS"
// 	corsAllowOrigin      = "*"
// 	corsAllowCredentials = "true"
// )

// MuxRouter for stock-catalog
func MuxRouter(api *routing.RouteGroup) {
	api.Post("/login", login.ExecuteCRUD())
	api.Get("/authed", isAuthed.ExecuteCRUD())
	// api.Post("/login", func(ctx *routing.Context) error {
	// 	ctx.Request.WriteTo(os.Stdout)
	// 	ctx.Response.Header.Set("Access-Control-Allow-Method", corsAllowMethods)
	// 	ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
	// 	ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
	// 	ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
	// 	return nil
	// })
}
