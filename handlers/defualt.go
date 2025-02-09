package handler

import (
	"encoding/json"
	"github.com/SigurdRiseth/CountryInfoService/utils"
	"net/http"
)

// DefaultHandler handles unknown routes and returns a JSON response
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// Create error response
	response := utils.ErrorResponse{
		Error:   http.StatusNotFound,
		Message: "You seem lost. The requested page was not found.",
	}

	// Encode response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
	}
}
