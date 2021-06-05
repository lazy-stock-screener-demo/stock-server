package readselfinfo

import (
	"encoding/json"
	appcore "stock-contexts/pkg/shared/application"

	"github.com/graphql-go/graphql"
	routing "github.com/qiangxue/fasthttp-routing"
)

type controller struct {
	restFul appcore.IBaseController
	useCase iuseCase
}

// Execute method
// fmt.Fprintf(ctx, "POST stockVID=%s <br/>", req)
// Get Header Value: ctx.Request.Header.Peek
// queryString := ctx.QueryArgs()
// fmt.Println("userName in read selfInfo", ctx.Request.Header.Peek("userName"))

func (c *controller) executeImpl(ctx *routing.Context) {
	reqDTO := ReqDTO{
		CustomerName: string(ctx.Request.Header.Peek("userName")),
	}
	customerDTO, eitherErr := c.useCase.execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "CustomerNotFound":
			c.restFul.NotFound(ctx, err.Message)
			return
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	} else {
		resDTO := ResDTO{customerDTO}
		dtoJSON, _ := json.Marshal(resDTO)
		c.restFul.OK(ctx, string(dtoJSON))
		return
	}
}

// NewController Construct new GetStockByTickerUseCase
func newController() *controller {
	return &controller{
		useCase: newUseCase(),
		restFul: appcore.NewBaseController(),
	}
}

type graphQLController struct {
	graphQL appcore.IGraphQLController
	useCase iuseCase
}

func (c *graphQLController) executeImpl(parms graphql.ResolveParams) (interface{}, interface{}) {
	reqDTO := ReqDTO{
		CustomerName: string(parms.Args["customerName"].(string)),
	}
	customerDTO, eitherErr := c.useCase.execute(reqDTO)

	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "TickerNotFound":
			return nil, err.Message
		default:
			return nil, err.Message
		}
	} else {
		resDTO := ResDTO{customerDTO}
		dtoJSON, _ := json.Marshal(resDTO)
		return []byte(string(dtoJSON)), nil
	}
}

func newGraphQLController() *graphQLController {
	return &graphQLController{
		useCase: newUseCase(),
		graphQL: appcore.NewGraphQLBaseController(),
	}
}
