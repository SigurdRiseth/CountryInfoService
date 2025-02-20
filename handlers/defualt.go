package handler

import (
	"net/http"
)

func DefaultHandler(writer http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(writer, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(writer, r, "../utils/index.html")
}
