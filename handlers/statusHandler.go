package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"math"
	"net/http"
	"time"
)

var StartTime time.Time

// GetStatus handles requests to the /status endpoint. It responds with the current
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
func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	countriesNowAPIStatus := getAPIStatus(COUNTRIES_NOW_API_URL)
	restCountriesAPIStatus := getAPIStatus(REST_COUNTRIES_API_URL)

	status := utils.APIStatus{
		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          API_VERSION,
		Uptime:           math.Round(time.Since(StartTime).Seconds()),
	}

	response, err := json.Marshal(status)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

// getAPIStatus checks if the given API is up and returns its status
func getAPIStatus(apiURL string) int { // TODO: 404 as valid status return or only 200s?
	// Send a GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error contacting API: %v", err)
		return http.StatusServiceUnavailable // Return 503 if there's an error
	}
	defer resp.Body.Close()

	// Check if the status code indicates success (200 OK)
	if resp.StatusCode == http.StatusOK {
		return http.StatusOK // Return 200 if the API is up
	}

	// If not 200 OK, log and return the received status code
	log.Printf("API responded with status: %v", resp.StatusCode)
	return resp.StatusCode // Return whatever status code the API sent
}
