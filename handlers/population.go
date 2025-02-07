package handler

import (
	"fmt"
	"net/http"
)

func GetPopulation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprint(w, `{"message": "This is the population endpoint"}`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
