package handler

import (
	"fmt"
	"net/http"
	"time"
)

// Declare startTime as a package-level variable
var startTime time.Time

// Initialize startTime when the package is initialized
func init() {
	startTime = time.Now()
}

// GetStatus returns the status of the service
func GetStatus(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, `{"status": "OK", "version": "v1", "uptime": "%s"}`, uptime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
