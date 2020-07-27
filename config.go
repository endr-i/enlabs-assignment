package main

import (
	"assignment/pg"
	"assignment/server"
)

type Config struct {
	Server server.Config
	DB     pg.Config
	Log    LogConfig
}
