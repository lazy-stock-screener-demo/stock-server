package middleware

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS middleware
func CORS() func(*routing.Context) error {
	return func(ctx *routing.Context) error {
		ctx.Response.Header.Set("Content-Type", "*/*")
		ctx.Response.Header.Set("Access-Control-Allow-Method", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		// ctx.Response.Header.WriteTo(os.Stdout)
		return nil
	}
}
