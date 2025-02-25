package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"log"
	"net/http"
	"time"
)

// Messages indicating the status of an API
const (
	OnlineMessage  = "Online"
	OfflineMessage = "Offline"
)

// StartTime stores the time when the server last started.
var StartTime time.Time

// HandleStatus retrieves the service status of external APIs and returns it as a JSON response.
//
// Parameters:
//   - w (http.ResponseWriter): The response writer to send the status data.
//   - r (*http.Request): The incoming HTTP request that triggers this handler.
//
// Returns:
//   - error: An error if an issue occurs during processing; otherwise, returns nil.
//
// Behavior:
//   - The function retrieves the status of the CountriesNow API and the RestCountries API by calling `getAPIStatus`.
//   - It then generates a response containing the statuses of both APIs along with the uptime of the service.
//   - The response is sent back to the client in JSON format with an HTTP status code of 200 (OK).
//
// Note:
//   - This function does not handle individual API errors beyond the request failure (e.g., network issues).
//   - The function relies on `getAPIStatus` to assess the availability of the external APIs and does not examine their HTTP status codes in detail (i.e., a 404 or 500 error is treated the same as a 200).
//   - It uses `NewAPIStatus` to create a structured response for the service status.
func HandleStatus(w http.ResponseWriter, r *http.Request) error {
	log.Println("Retrieving service status")
	w.Header().Set("Content-Type", "application/json")

	// Fetch statuses of external APIs
	apiStatuses := utils.NewAPIStatus(
		getAPIStatus(utils.CountriesNowApiUrl),
		getAPIStatus(utils.RestCountriesApiUrl),
		time.Since(StartTime).Seconds(),
	)

	// Create API response
	resp := utils.APIResponse{
		Error:   false,
		Message: "Service status retrieved successfully",
		Data:    apiStatuses,
	}

	// Encode response to JSON
	response, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	return err
}

// getAPIStatus checks the availability of an API by sending a GET request to the provided URL.
//
// Parameters:
//   - apiURL (string): The URL of the API to check.
//
// Returns:
//   - string: A predefined message indicating whether the API is online or offline.
//
// Behavior:
//   - If the request fails (e.g., network error, timeout), it logs the error and returns `OfflineMessage`.
//   - If the request succeeds (regardless of status code), it returns `OnlineMessage`.
//
// Note:
//   - This function does not differentiate between different HTTP status codes (e.g., 404 vs. 200).
//   - The API is considered "online" as long as the server responds to the request.
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
