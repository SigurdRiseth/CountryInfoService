package handler

import (
	"encoding/json"
	utils2 "github.com/SigurdRiseth/CountryInfoService/internal/utils"
	"log"
	"net/http"
	"time"
)

const (
	OnlineMessage  = "Online"
	OfflineMessage = "Offline"
)

var StartTime time.Time

// HandleStatus handles requests to the /status endpoint. It responds with the current
// status of the API, including:
// - HTTP status codes for each service (CountriesNowAPI and RestCountriesAPI)
// - The current version of the API
// - The service's uptime in seconds
//
// The function sets the appropriate Content-Type header and responds with a
// JSON object containing the service's status. If any error occurs during
// the encoding or writing process, an HTTP 500 error is returned.
//
// It does not take any parameters directly and returns the status in JSON format.
func HandleStatus(w http.ResponseWriter, r *http.Request) error {
	log.Println("Retrieving service status")
	w.Header().Set("Content-Type", "application/json")

	countriesNowAPIStatus := getAPIStatus(utils2.CountriesNowApiUrl)
	restCountriesAPIStatus := getAPIStatus(utils2.RestCountriesApiUrl)

	status := utils2.NewAPIStatus(
		countriesNowAPIStatus,
		restCountriesAPIStatus,
		time.Since(StartTime).Seconds(),
	)

	resp := utils2.APIResponse{
		Error:   false,
		Message: "Service status retrieved successfully",
		Data:    status,
	}

	response, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		return err
	}

	return nil
}

// getAPIStatus checks the status of the given API URL by sending a GET request.
// If the request is successful, it returns "Online". If there is an error
// during the request, it returns "Offline".
//
// Parameters:
// - apiURL: The URL of the API to check.
//
// Returns:
// - A string indicating the status of the API ("Online" or "Offline").
func getAPIStatus(apiURL string) string {
	// Send a GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return OfflineMessage
	}
	defer resp.Body.Close()

	return OnlineMessage
}
