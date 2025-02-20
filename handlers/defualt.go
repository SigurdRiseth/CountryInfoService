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

	// Add a list of endpoints and how to call them, including query parameters
	endpoints := []string{
		"The following endpoints are available:",
		"GET /country/v1/info/{countryCode} - Get information about a country. Example: /country/v1/info/NO",
		"  Optional query parameter: ?limit=x (limits number of cities). Example: /country/v1/info/NO?limit=5",
		"GET /country/v1/population/{countryCode} - Get population data. Example: /country/v1/population/US",
		"  Optional query parameter: ?limit=YYYY-YYYY (limits data to a year range). Example: /country/v1/population/US?limit=2002-2008",
		"GET /country/v1/status - Check the service status.",
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
