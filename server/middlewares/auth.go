package middlewares

import "github.com/valyala/fasthttp"

func AuthMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		// TODO: auth
		handler(ctx)
	}
}
