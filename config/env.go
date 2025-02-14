package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnvVariables loads the environment variables from the .env file
func loadEnvVariables() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Printf("Error loading environment variables: %v", err)
	}
}

// GetPort retrieves the port from the environment variable or defaults to 8080
func GetPort() string {

	loadEnvVariables()

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Defaulting to 8080")
		port = "8080"
	}
	return port
}
