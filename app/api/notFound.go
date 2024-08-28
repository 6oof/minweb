package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Create the error response
	errorResponse := ErrorResponse{
		Error:  "Resource not found",
		Status: 404,
	}

	// Set the content-type header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// Write the JSON error response
	json.NewEncoder(w).Encode(errorResponse)
}
