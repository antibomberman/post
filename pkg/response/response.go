package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	response := response{
		Message: message,
		Data:    data,
	}
	sendResponse(w, response, http.StatusOK)
}

func Error(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	response := response{
		Message: message,
		Data:    data,
	}
	sendResponse(w, response, statusCode)
}

func sendResponse(w http.ResponseWriter, response response, statusCode int) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
