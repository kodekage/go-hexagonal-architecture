package main

import (
	"github.com/kodekage/banking/cmd"
	"github.com/kodekage/banking/internal/logger"
)

func main() {
	logger.Info("Starting application")
	cmd.StartServer()
}
