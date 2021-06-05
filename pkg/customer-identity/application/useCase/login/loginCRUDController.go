package login

import (
	"encoding/json"
	identitymapper "stock-contexts/pkg/customer-identity/application/mapper/Identity"
	appcore "stock-contexts/pkg/shared/application"

	routing "github.com/qiangxue/fasthttp-routing"
)

type controllerCRUD struct {
	restFul appcore.IBaseController
	useCase IUseCaseCRUD
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

	loginView, eitherErr := c.useCase.Execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "UserIDNotFound":
			c.restFul.NotFound(ctx, err.Message)
			return
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	} else {
		resDTO := ResDTO{identitymapper.NewEntityMap().ToDTO(*loginView)}
		dtoJSON, _ := json.Marshal(resDTO)
		// ctx.Response.Header.Set("Authorization", "Bearer "+resDTO.Identity.AccessToken)
		// ctx.Response.Header.Set("Authorization", "Bearer "+resDTO.Identity.AccessToken)
		c.restFul.OK(ctx, string(dtoJSON))
		return
	}
}

// NewController Construct new GetStockByTickerUseCase
func newControllerCRUD() *controllerCRUD {
	return &controllerCRUD{
		useCase: NewUseCaseCRUD(),
		restFul: appcore.NewBaseController(),
	}
}
