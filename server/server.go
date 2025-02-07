package server

import (
	"github.com/SigurdRiseth/CountryInfoService/config"
	"log"
	"net/http"
	"time"

	handler "github.com/SigurdRiseth/CountryInfoService/handlers"
	"github.com/SigurdRiseth/CountryInfoService/utils"
)

// StartServer initializes and starts the HTTP server
func StartServer() {
	handler.StartTime = time.Now() // Initialize start time

	// Load environment variables
	config.LoadEnvVariables()

	// Get the port from environment variables, default to 8080
	port := config.GetPort()

	// Instantiate the router
	router := setupRouter()

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router)) // TODO: Gracefully shutdown server?
}

// setupRouter sets up the HTTP routes and handlers
func setupRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Define the endpoints
	router.HandleFunc(utils.INFO_PATH, handler.GetInfo)
	router.HandleFunc(utils.POPULATION_PATH, handler.GetPopulation)
	router.HandleFunc(utils.STATUS_PATH, handler.GetStatus)
	router.HandleFunc("/", handler.DefaultHandler)

	return router
}
