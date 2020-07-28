package utils

import (
	"encoding/json"
	"errors"
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

type Message struct {
	ErrorMessage StringNull
	Data         interface{}
}

func NewResponse(data interface{}, err error, logger *log.Entry) IResponse {
	return &Response{
		Error: err,
		Data:  data,
		Log:   logger,
	}
}

func (r Response) message() Message {
	var errorMessage string
	if r.Error != nil {
		errorMessage = r.Error.Error()
	}
	return Message{
		ErrorMessage: StringNull{
			IsNull: r.Error == nil,
			Value:  errorMessage,
		},
		Data: r.Data,
	}
}

func (r Response) Print(ctx *fasthttp.RequestCtx) {
	if r.Error != nil {
		r.Log.WithError(r.Error).Warn()
		res, _ := json.Marshal(r.message())
		ctx.Error(string(res), getStatus(ctx, fasthttp.StatusBadRequest))
	} else if msg, err := json.Marshal(r.message()); err != nil {
		r.Log.WithError(err).Warn()
		r.Error = errors.New("cannot convert response")
		res, _ := json.Marshal(r.message())
		ctx.Error(string(res), fasthttp.StatusInternalServerError)
	} else {
		r.Log.Info(r.Data)
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
