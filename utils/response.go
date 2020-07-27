package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type IResponse interface {
	Print(ctx *fasthttp.RequestCtx)
}

type Response struct {
	Error error
	Data  interface{}
	Log   *log.Entry
}

func NewResponse(data interface{}, err error, logger *log.Entry) IResponse {
	return &Response{
		Error: err,
		Data:  data,
		Log:   logger,
	}
}

func (r Response) Print(ctx *fasthttp.RequestCtx) {
	if r.Error != nil {
		r.Log.WithError(r.Error).Warn()
		ctx.Error(r.Error.Error(), getStatus(ctx, fasthttp.StatusBadRequest))
	} else if msg, err := json.Marshal(r.Data); err != nil {
		r.Log.WithError(err).Warn()
		ctx.Error("cannot convert response", fasthttp.StatusInternalServerError)
	} else {
		r.Log.Info(msg)
		ctx.Success("application/json", msg)
	}
}

func getStatus(ctx *fasthttp.RequestCtx, defaultStatus int) int {
	status := ctx.Response.Header.StatusCode()
	if status != fasthttp.StatusOK {
		return status
	}
	return defaultStatus
}
