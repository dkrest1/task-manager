package utils

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Payload   interface{} `json:"payload"`

}

func NewGenericResponse(status int, message string, payload interface{}) *GenericResponse {
	return &GenericResponse{
		Status: status,
		Message: message,
		Payload: payload,
	}
}

func HandleGenericResponse(w http.ResponseWriter, message string, statusCode int) {
	response := NewGenericResponse(statusCode, message, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}