package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"net/http"
)

// DefaultHandler handles unknown routes and returns a JSON response
func DefaultHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// Add a list of endpoints and how to call them
	endpoints := []string{
		"The following endpoints are available:",
		"GET /countryinfo/v1/info/{countryCode} - Get information about a country. Example: /countryinfo/v1/info/NO",
		"GET /countryinfo/v1/population/{countryCode} - Get population data. Example: /countryinfo/v1/population/US",
		"GET /countryinfo/v1/status - Check the service status.",
	}

	// Create error response
	response := utils.APIResponse{
		Error:   true,
		Message: "You seem lost. The requested page was not found.",
		Data:    endpoints,
	}

	// Encode response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}
