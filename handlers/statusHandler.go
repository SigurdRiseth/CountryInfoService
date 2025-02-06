package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"math"
	"net/http"
	"time"
)

var StartTime time.Time

const apiVersion = "v1"

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

	countriesNowAPIStatus := getCountriesNowAPIStatus()
	restCountriesAPIStatus := getRestCountriesAPIStatus()

	status := utils.APIStatus{
		CountriesNowAPI:  restCountriesAPIStatus,
		RestCountriesAPI: countriesNowAPIStatus,
		Version:          apiVersion,
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

// getCountriesNowAPIStatus returns the status of the CountriesNowAPI.
func getCountriesNowAPIStatus() int {
	// Placeholder for CountriesNowAPI status
	return http.StatusOK
}

// getRestCountriesAPIStatus returns the status of the RestCountriesAPI.
func getRestCountriesAPIStatus() int {
	// Placeholder for RestCountriesAPI status
	return http.StatusOK
}
