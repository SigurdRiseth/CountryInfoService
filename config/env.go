package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// loadEnvVariables loads environment variables from a .env file located in the parent directory.
// If there is an error loading the .env file, it logs the error.
func loadEnvVariables() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading environment variables: %v", err)
	}
}

// GetPort retrieves the port from the environment variable "PORT" or defaults to 8080 if the variable is not set.
//
// Returns:
// - A string representing the port number.
func GetPort() string {
	// Load .env only if running locally
	if os.Getenv("RENDER") == "" { // Render automatically sets this variable
		log.Println("Running locally. Loading environment variables from .env file")
		loadEnvVariables()
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Defaulting to 8080")
		port = "8080"
	}
	return port
}
