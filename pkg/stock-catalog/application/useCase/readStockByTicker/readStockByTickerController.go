package readstockbyticker

import (
	"encoding/json"
	appcore "stock-contexts/pkg/shared/application"
	catalogmapper "stock-contexts/pkg/stock-catalog/application/mapper/Catalog"

	"github.com/graphql-go/graphql"
	routing "github.com/qiangxue/fasthttp-routing"
)

type controller struct {
	restFul appcore.IBaseController
	useCase iuseCase
}

// Execute method
func (c *controller) executeImpl(ctx *routing.Context) {
	queryString := ctx.QueryArgs()
	reqDTO := ReqDTO{
		StockVID: string(queryString.Peek("stockVID")),
	}
	stockView, eitherErr := c.useCase.execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "TickerNotFound":
			c.restFul.NotFound(ctx, err.Message)
			return
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	} else {
		// fmt.Println("stockView", stockView.GetStockVID().GetValue())
		resDTO := ResDTO{catalogmapper.NewViewMap().ToDTO(stockView)}
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
		StockVID: string(parms.Args["customerVID"].(string)),
	}
	stockView, eitherErr := c.useCase.execute(reqDTO)

	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "TickerNotFound":
			return nil, err.Message
		default:
			return nil, err.Message
		}
	} else {
		// fmt.Println(eitherErr.Result.GetData())
		// resDTO := NewResDTO(catalogmapper.NewViewMap().ToDTO(stockView).GetDTO())
		resDTO := ResDTO{catalogmapper.NewViewMap().ToDTO(stockView)}
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
