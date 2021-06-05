package stockcataloggraphql

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	routing "github.com/qiangxue/fasthttp-routing"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"stock": &graphql.Field{
				Type:        stockType,
				Description: "Read stock by ticker",
				Args: graphql.FieldConfigArgument{
					"vid": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: readStockByTicker,
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

// Handler define a method to build
func Handler() func(ctx *routing.Context) error {
	return func(ctx *routing.Context) error {
		// r.URL.Query().Get("query")
		print(ctx.Request.URI().QueryString())
		result := executeQuery(string(ctx.Request.URI().QueryString()), schema)
		json.NewEncoder(ctx.Response.BodyWriter()).Encode(result)
		return nil
	}
}

// func TestHttpOnly() {
// 	http.HandleFunc("/graphql/v1", func(w http.ResponseWriter, r *http.Request) {
// 		result := executeQuery(r.URL.Query().Get("query"), schema)
// 		json.NewEncoder(w).Encode(result)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
