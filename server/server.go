package server

import (
	"assignment/server/middlewares"
	"assignment/server/routes/ping"
	"assignment/server/routes/user"
	"assignment/server/routes/user/transaction"
	"assignment/utils"
	"encoding/json"
	"github.com/buaazp/fasthttprouter"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Config struct {
	Host    string `default:":8080"`
	LogFile string `default:""`
}

func Router() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/ping", ping.HandleGet())
	router.POST("/user", middlewares.AuthMiddleware(user.HandlePost()))
	router.GET("/user/:userId", middlewares.AuthMiddleware(user.HandleGet()))
	router.POST("/user/:userId/transaction", middlewares.AuthMiddleware(transaction.HandlePost()))
	router.PanicHandler = handlePanic
	return router
}

func handlePanic(ctx *fasthttp.RequestCtx, i interface{}) {
	log.WithFields(map[string]interface{}{
		"url":     ctx.URI().String(),
		"body":    string(ctx.Request.Body()),
		"headers": ctx.Request.Header.String(),
		"panic":   spew.Sdump(i),
	}).Error("panic recovered")
	message := utils.Message{
		ErrorMessage: utils.StringNull{Value: "internal server error"},
	}
	msg, _ := json.Marshal(message)
	ctx.Error(string(msg), fasthttp.StatusInternalServerError)
}
