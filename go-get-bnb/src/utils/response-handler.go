package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
}

func JSONResponse(w http.ResponseWriter, status string, statusCode int, data interface{}, errors interface{}) {
	response := APIResponse{
		Status:     status,
		StatusCode: statusCode,
		Data:       data,
		Error:      errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
