package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"net/http"
)

// AvailableEndpoints defines the API endpoints and how to use them
var AvailableEndpoints = []map[string]string{
	{"endpoint": "GET /country/v1/info/{countryCode}", "description": "Get information about a country.", "example": "/country/v1/info/NO"},
	{"endpoint": "Optional query parameter", "description": "?limit=x (limits number of cities).", "example": "/country/v1/info/NO?limit=5"},
	{"endpoint": "GET /country/v1/population/{countryCode}", "description": "Get population data.", "example": "/country/v1/population/US"},
	{"endpoint": "Optional query parameter", "description": "?limit=YYYY-YYYY (limits data to a year range).", "example": "/country/v1/population/US?limit=2002-2008"},
	{"endpoint": "GET /country/v1/status", "description": "Check the service status.", "example": "/country/v1/status"},
}

// DefaultHandler handles unknown routes and returns a JSON response
func DefaultHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// Create error response
	response := utils.APIResponse{
		Error:   true,
		Message: "You seem lost. The requested page was not found.",
		Data:    AvailableEndpoints,
	}

	// Encode response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}
