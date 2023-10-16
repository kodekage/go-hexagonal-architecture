package cmd

import (
	"log"
	"net/http"
)

func StartServer() {
	router := setupRoutes()

	// Start cmd
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
