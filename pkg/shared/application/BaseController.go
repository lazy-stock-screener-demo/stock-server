package appcore

import (
	"stock-contexts/pkg/shared/config"

	routing "github.com/qiangxue/fasthttp-routing"
)

// IBaseController define base interface
type IBaseController interface {
	OK(ctx *routing.Context, dto string)
	Created(ctx *routing.Context, message string)
	ClientError(ctx *routing.Context, message string)
	Unauthorized(ctx *routing.Context, message string)
	NotFound(ctx *routing.Context, message string)
	Conflict(ctx *routing.Context, message string)
	TooMany(ctx *routing.Context, message string)
	Fail(ctx *routing.Context, Error interface{})
}

// BaseController define a struct
type BaseController struct{}

// OK define 200
func (b *BaseController) OK(ctx *routing.Context, dto string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(200)
	ctx.SetBody([]byte(dto))
}

// Created define 201
func (b *BaseController) Created(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(201)
	ctx.SetBody([]byte(message))
}

// ClientError define 400
func (b *BaseController) ClientError(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(400)
	ctx.SetBody([]byte(message))
}

// Unauthorized define 401
func (b *BaseController) Unauthorized(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(401)
	ctx.SetBody([]byte(message))
}

// NotFound define 404
func (b *BaseController) NotFound(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(404)
	ctx.SetBody([]byte(message))
}

// Conflict define 409
func (b *BaseController) Conflict(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(409)
	ctx.SetBody([]byte(message))
}

// TooMany define 429
func (b *BaseController) TooMany(ctx *routing.Context, message string) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(429)
	ctx.SetBody([]byte(message))
}

// Fail defin 500
func (b *BaseController) Fail(ctx *routing.Context, Error interface{}) {
	ctx.SetContentType(config.AcceptJSON)
	ctx.SetStatusCode(500)
	switch Error.(type) {
	case string:
		ctx.SetBody([]byte(Error.(string)))
	default:
		ctx.SetBody([]byte(Error.([]byte)))
	}
}

// NewBaseController construct new base constroller
func NewBaseController() *BaseController {
	return &BaseController{}
}
