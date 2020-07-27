package main

import (
	"assignment/models"
	"assignment/pg"
	"assignment/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/configor"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	var config Config
	configor.New(&configor.Config{
		ENVPrefix: "ENLABS",
	}).Load(&config)
	spew.Dump(config)
	logInit(config.Log)
	router := server.Router()
	if db, err := pg.Init(config.DB, models.Migrations...); err != nil {
		log.Fatal(err)
	} else {
		models.InitRepositories(db)
	}
	log.Fatal(fasthttp.ListenAndServe(config.Server.Host, router.Handler))
}
