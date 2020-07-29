package ping

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func HandleGet() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "pong")
	}
}
