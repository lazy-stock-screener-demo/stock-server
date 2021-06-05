package readautocompletebychar

import (
	"encoding/json"
	appcore "stock-contexts/pkg/shared/application"

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
		StockChar: string(queryString.Peek("char")),
	}
	matchingList, eitherErr := c.useCase.execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	} else {
		// fmt.Println("list", list.GetStockVID().GetValue())
		resDTO := ResDTO{matchingList}
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
