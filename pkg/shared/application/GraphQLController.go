package appcore

// IGraphQLController define base interface
type IGraphQLController interface{}

// GraphQLController define a struct
type GraphQLController struct{}

// NewGraphQLBaseController define a graphql base controller
func NewGraphQLBaseController() *GraphQLController {
	return &GraphQLController{}
}
