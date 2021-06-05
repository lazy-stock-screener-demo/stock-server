package isauthed

import (
	"fmt"
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
	// queryString := ctx.QueryArgs()
	reqDTO := ReqDTO{
		AuthorizationHeader: string(ctx.Request.Header.Peek("Authorization")),
	}

	userName, eitherErr := c.useCase.execute(reqDTO)
	if eitherErr.IsError() {
		err := eitherErr.Result.GetErr()
		switch errType := err.Type; errType {
		case "UserNotFound", "TokenSignatureExpired", "NoAccessToken":
			ctx.Response.Header.Set("WWW-Authenticate", fmt.Sprintf(`Bearer realm="%s"`, "lazy-stock-screener.com"))
			c.restFul.Unauthorized(ctx, err.Message)
			ctx.Abort()
			return
		default:
			c.restFul.Fail(ctx, err.Error)
			return
		}
	}
	fmt.Println(userName)
	ctx.Request.Header.Set("userName", userName)
	// No error just pass to next
}

// NewController Construct new GetStockByTickerUseCase
func newControllerCRUD() *controllerCRUD {
	return &controllerCRUD{
		useCase: NewUseCaseCRUD(),
		restFul: appcore.NewBaseController(),
	}
}
