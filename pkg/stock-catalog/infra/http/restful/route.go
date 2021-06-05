package catalogrestful

import (
	readAutoCompleteByChar "stock-contexts/pkg/stock-catalog/application/useCase/readAutoCompleteByChar"
	readstockbyticker "stock-contexts/pkg/stock-catalog/application/useCase/readStockByTicker"

	routing "github.com/qiangxue/fasthttp-routing"
)

// MuxRouter for stock-catalog
func MuxRouter(api *routing.RouteGroup) {
	api.Get("/stock", readstockbyticker.Execute())
	api.Get("/search", readAutoCompleteByChar.Execute())
}

// r.GET("/", Index)
// r.GET("/hello/{name}", Hello)
// func Routes(router *mux.Router) *mux.Router {
// 	loginRouter := router.PathPrefix("/login").Subrouter()
// 	restful.UserRouter(loginRouter)
// 	return router
// }
