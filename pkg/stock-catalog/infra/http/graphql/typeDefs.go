package stockcataloggraphql

import "github.com/graphql-go/graphql"

var stockType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
