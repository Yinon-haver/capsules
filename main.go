package main

import (
	"github.com/capsules-web-server/config"
	"github.com/capsules-web-server/db"
	"github.com/capsules-web-server/logger"
	"github.com/capsules-web-server/server"
)


func main() {
	logger.Init()
	config.Init()
	db.Init()
	server.Run()
}