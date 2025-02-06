package main

import (
	handler "REST-stub/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	// Load the environment variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	// Extract the port from the environment variables
	port := os.Getenv("PORT")

	// If the port is not set, default to 8080
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Instantiate the router
	router := http.NewServeMux()

	// Define the endpoints
	router.HandleFunc(handler.INFO_PATH, handler.GetInfo)
	router.HandleFunc(handler.POPULATION_PATH, handler.GetPopulation)
	router.HandleFunc(handler.STATUS_PATH, handler.GetStatus)

	router.HandleFunc("/", handler.HomePage)

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
