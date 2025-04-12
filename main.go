package main

import (
	"log"
	"net/http"
	"os"

	"github.com/greedy_game/targeting_engine/service"
	"github.com/greedy_game/targeting_engine/transport"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Create service
	var log log.Logger
	svc := service.NewService(&log)

	// Create HTTP handler
	handler := transport.NewHTTPHandler(svc)

	// Start server
	logger.Print("Starting server on :8080")
	logger.Fatal(http.ListenAndServe(":8080", handler))
}
