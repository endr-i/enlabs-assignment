package main

import (
	"assignment/entities/migrations"
	"assignment/entities/repos"
	"assignment/entities/repos/common"
	"assignment/pg"
	"assignment/server"
	"assignment/server/schedulers"
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
	db, err := pg.Init(config.DB, migrations.Migrations...)
	if err != nil {
		log.Fatal(err)
	}
	repos.InitRepositories(db)
	if config.Schedulers.Enable {
		schedulers.InitSchedulers(config.Schedulers.CancelOddConfig, common.GetRepo())
	}
	log.Fatal(fasthttp.ListenAndServe(config.Server.Host, router.Handler))
}
