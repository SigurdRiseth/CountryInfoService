package handler

import (
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound) // Set HTTP status code to 404

	_, err := fmt.Fprint(w, "<h1>You seem lost</h1><p>404 - Page Not Found</p>")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
