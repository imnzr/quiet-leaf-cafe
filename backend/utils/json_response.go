package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonError(writter http.ResponseWriter, statusCode int, data interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(statusCode)
	err := json.NewEncoder(writter).Encode(data)
	if err != nil {
		http.Error(writter, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
