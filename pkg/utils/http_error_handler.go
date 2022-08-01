package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"rollic/internal/models"
)

func HTTPErrorHandler(response http.ResponseWriter, error string, status int) {
	response.WriteHeader(status)
	errorResponse := models.HttpErrorResponse{
		Error: error,
	}
	jData, _ := json.Marshal(errorResponse)
	_, err := response.Write(jData)
	if err != nil {
		log.Fatalf("Error writing response: %v", err)
	}
}
