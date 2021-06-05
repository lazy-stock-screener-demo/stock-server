package graphqlapi

import (
	stockcataloggraphql "stock-contexts/pkg/stock-catalog/infra/http/graphql"

	routing "github.com/qiangxue/fasthttp-routing"
)

// V1Router build graphql router
func V1Router(router *routing.Router) *routing.Router {
	router.Any("/graphql/v1", stockcataloggraphql.Handler())
	return router
}

// func TestHttpOnly() {
// 	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
// 		result := executeQuery(r.URL.Query().Get("query"), schema)
// 		json.NewEncoder(w).Encode(result)
// 	})
// }
