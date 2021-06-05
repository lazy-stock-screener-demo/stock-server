package createcustomer

import (
	"encoding/json"
	appcore "stock-contexts/pkg/shared/application"

	routing "github.com/qiangxue/fasthttp-routing"
)

type controllerCRUD struct {
	restFul appcore.IBaseController
	useCase iuseCaseCRUD
}

// Execute method
func (c *controllerCRUD) executeImpl(ctx *routing.Context) {
	// ctx.Request.WriteTo(os.Stdout)
	var reqDTO ReqDTO
	postValues := ctx.PostBody()
	if err := json.Unmarshal(postValues, &reqDTO); err != nil {
		c.restFul.Fail(ctx, err.Error())
		return
	}

	customerDTO, eitherErr := c.useCase.execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "CustomerSavingInRepo":
			c.restFul.Fail(ctx, err.Message)
			return
		case "LicenseNotFound":
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
func newControllerCRUD() *controllerCRUD {
	return &controllerCRUD{
		useCase: newUseCaseCRUD(),
		restFul: appcore.NewBaseController(),
	}
}
