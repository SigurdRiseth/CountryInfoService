package main

import (
	handler "github.com/SigurdRiseth/CountryInfoService/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	handler.StartTime = time.Now() // Initialize start time

	// Load environment variables
	if err := loadEnvVariables(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Get the port from environment variables, default to 8080
	port := getPort()

	// Instantiate the router
	router := setupRouter()

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// loadEnvVariables loads the environment variables from the .env file
func loadEnvVariables() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

// getPort retrieves the port from the environment variable or defaults to 8080
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Defaulting to 8080")
		port = "8080"
	}
	return port
}

// setupRouter sets up the HTTP routes and handlers
func setupRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Define the endpoints
	router.HandleFunc(handler.INFO_PATH, handler.GetInfo)
	router.HandleFunc(handler.POPULATION_PATH, handler.GetPopulation)
	router.HandleFunc(handler.STATUS_PATH, handler.GetStatus)
	router.HandleFunc("/", handler.HomePage)

	return router
}
