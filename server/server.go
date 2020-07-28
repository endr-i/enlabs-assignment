package server

import (
	"assignment/server/middlewares"
	"assignment/server/routes/user"
	"assignment/server/routes/user/transaction"
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
	router.POST("/user", middlewares.AuthMiddleware(user.HandlePost()))
	router.GET("/user/:userId", middlewares.AuthMiddleware(user.HandleGet()))
	router.POST("/user/:userId/transaction", middlewares.AuthMiddleware(transaction.HandlePost()))
	return router
}
