package utils

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]string{"message": message}
	SendJSONResponse(w, statusCode, response)
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
