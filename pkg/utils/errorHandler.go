package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func ErrorHandler(w http.ResponseWriter, internalErr error, clientMsg string, statusCode int) {

	if internalErr != nil {
		errorLogger.Printf("%s: %v", clientMsg, internalErr)
	} else {
		errorLogger.Printf("%s", clientMsg)
	}

	// Send friendly message to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := ErrorResponse{
		Success: false,
		Message: clientMsg,
	}

	json.NewEncoder(w).Encode(resp)

}
