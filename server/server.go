package server

import (
	"github.com/SigurdRiseth/CountryInfoService/config"
	"github.com/SigurdRiseth/CountryInfoService/handlers"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
	"time"
)

// StartServer initializes and starts the HTTP server.
func StartServer() {
	handler.StartTime = time.Now() // Initialize start time

	// Get the port from environment variables, default to 8080
	port := config.GetPort()

	// Instantiate the router
	router := setupRouter()

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// setupRouter initializes the HTTP request multiplexer (router) and defines the endpoints.
//
// Returns:
// - *http.ServeMux: A pointer to the initialized HTTP request multiplexer.
func setupRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Define the endpoints
	router.HandleFunc(utils.GetInfoPath(""), makeHTTPHandleFunc(handler.HandleInfo))
	router.HandleFunc(utils.GetPopulationPath(""), makeHTTPHandleFunc(handler.HandlePopulation))
	router.HandleFunc(utils.GetStatusPath(), makeHTTPHandleFunc(handler.HandleStatus))
	router.HandleFunc("/", handler.DefaultHandler)

	return router
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

// makeHTTPHandleFunc is a helper function that wraps an apiFunc with error handling.
// It returns an http.HandlerFunc that logs the error and sends an HTTP 500 status code
// if the wrapped function returns an error.
//
// Parameters:
// - f: The apiFunc to be wrapped.
//
// Returns:
// - http.HandlerFunc: A function that handles HTTP requests and responses.
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// TODO: switch on error type and return appropriate status code
		}
	}
}
