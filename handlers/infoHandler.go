package handler

import (
	"fmt"
	"net/http"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprint(w, `{"message": "This is the info endpoint"}`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
