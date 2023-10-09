package main

import (
	"github.com/kodekage/banking/app"
	"github.com/kodekage/banking/logger"
)

func main() {
	logger.Info("Starting application")
	app.StartServer()
}
