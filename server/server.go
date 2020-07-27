package server

import (
	"assignment/server/middlewares"
	"assignment/server/routes/transactions"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Config struct {
	Host    string `default:":8080"`
	LogFile string `default:""`
}

func Router() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "pong")
	})
	router.POST("/user/:userId/transaction", middlewares.AuthMiddleware(transactions.Handle()))
	return router
}
