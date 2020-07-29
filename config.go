package main

import (
	"assignment/pg"
	"assignment/server"
	cancel_odd "assignment/server/schedulers/cancel-odd"
)

type Config struct {
	Server     server.Config
	DB         pg.Config
	Log        LogConfig
	Schedulers struct {
		Enable          bool `default:"false"`
		CancelOddConfig cancel_odd.Config
	}
}
