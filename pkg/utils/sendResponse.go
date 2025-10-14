package utils

import (
	"encoding/json"
	"net/http"
)

type response[T any] struct {
	Message string `json:"message"`
	Data    *T     `json:"data,omitempty"`
	Success bool
}

func SendResponse[T any](w http.ResponseWriter, r *http.Request, message string, statusCode int, data *T) error {
	w.Header().Set("Content-Type", "application/json")
	res := &response[T]{
		Message: message,
		Data:    data,
		Success: true,
	}
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(res)
}
