package restfulapi

import (
	identityrestful "stock-contexts/pkg/customer-identity/infra/http/restful"
	customerrestful "stock-contexts/pkg/customer-self/infra/http/restful"
	catalogrestful "stock-contexts/pkg/stock-catalog/infra/http/restful"

	routing "github.com/qiangxue/fasthttp-routing"
)

// V1Router build version 1 router
func V1Router(router *routing.Router) *routing.Router {
	api := router.Group("/v1")
	catalogrestful.MuxRouter(api)
	identityrestful.MuxRouter(api)
	customerrestful.MuxRouter(api)
	// api.Get("/stock", func(c *routing.Context) error {
	// 	fmt.Println("Test Route")
	// 	return nil
	// })
	return router
}

// func V1Router() *mux.Router {
// 	router := mux.NewRouter()
// 	v1Router := router.PathPrefix("/v1").Subrouter()
// 	return restful.Routes(v1Router)
// }
