package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	// default status code
	defaultStatusCode := http.StatusInternalServerError
	// check if status code is valid
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	// response
	body := errorResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	w.WriteHeader(defaultStatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func Errorf(w http.ResponseWriter, statusCode int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	Error(w, statusCode, message)
}