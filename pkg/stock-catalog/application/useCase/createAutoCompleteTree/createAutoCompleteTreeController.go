package autocompletesearch

import (
	"encoding/json"
	appcore "stock-contexts/pkg/shared/application"
	ternarySearchTree "stock-contexts/pkg/shared/services/ternarySearchTree"

	routing "github.com/qiangxue/fasthttp-routing"
)

type controller struct {
	restFul appcore.IBaseController
	useCase iuseCase
}

// Execute method
func (c *controller) executeImpl(ctx *routing.Context) {
	// queryString := ctx.QueryArgs()
	// reqDTO := ReqDTO{
	// 	StockVID: string(queryString.Peek("stockVID")),
	// }
	_, eitherErr := c.useCase.execute()
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	} else {
		// fmt.Println("list", list.GetStockVID().GetValue())
		resDTO := ResDTO{"Creating Tree Node"}
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

type workerController struct {
	useCase iuseCase
}

func (c *workerController) executeImpl() *ternarySearchTree.Tree {
	TSTService, _ := c.useCase.execute()
	// if eitherErr.IsError() {
	// 	err := eitherErr.Result.GetErr()

	// } else {

	// }
	return TSTService
}

func newWorkerController() *workerController {
	return &workerController{
		useCase: newUseCase(),
	}
}
