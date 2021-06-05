package sendEmail

import (
	appcore "stock-contexts/pkg/shared/application"

	routing "github.com/qiangxue/fasthttp-routing"
)

type controller struct {
	restFul appcore.IBaseController
	useCase iuseCase
}

func (c *controller) executeImpl(ctx *routing.Context) {
	reqDTO := ReqDTO{
		Email: string(ctx.Request.Header.Peek("userName")),
		Link:  string(ctx.Request.Header.Peek("activate")),
	}
	_, eitherErr := c.useCase.execute(reqDTO)
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
		// resDTO := ResDTO{customerDTO}
		// dtoJSON, _ := json.Marshal(resDTO)
		// c.restFul.OK(ctx, string(dtoJSON))
		return
	}
}

func newController() *controller {
	return &controller{
		useCase: newUseCase(),
		restFul: appcore.NewBaseController(),
	}
}
