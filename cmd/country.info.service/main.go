package main

import (
	handler "REST-stub/handlers"
	"fmt"
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
	router.HandleFunc(handler.INFO_PATH, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := fmt.Fprint(w, `{"message": "This is the info endpoint"}`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	//router.HandleFunc(handler.POPULATION_PATH, GetPopulation)
	//router.HandleFunc(handler.STATUS_PATH, GetStatus)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := fmt.Fprint(w, `
			<!DOCTYPE html>
			<html>
			<head><title>Country Info API</title></head>
			<body>
				<h1>Welcome to the Country Info API</h1>
				<p>This API provides information about different countries.</p>
				<h2>Endpoints</h2>
				<ul>
					<li><code>/country/v1/info</code> - List country information</li>
					<li><code>/country/v1/population</code> - Get details about a countries population over time</li>
					<li><code>/country/v1/status</code> - Get details about the APIs status</li>
				</ul>
			</body>
			</html>
		`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the server
	log.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
