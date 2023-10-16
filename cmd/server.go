package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kodekage/banking/internal/logger"
)

func StartServer() {
	// setup environments and variables.
	envSetup()

	// setup server routes
	router := setupRoutes()

	// Start cmd
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), router))
}

func envSetup() {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("error loading .env file: " + err.Error())
	}
}
