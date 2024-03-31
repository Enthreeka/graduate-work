package handler

import (
	"encoding/json"
	"net/http"
)

type JSONError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func ErrorJSON(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError := &JSONError{
		Status: http.StatusText(code),
		Error:  error,
	}
	e := json.NewEncoder(w)
	e.Encode(jsonError)
}

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	ErrorJSON(w, err.Error(), statusCode)
}

func DecodingError(w http.ResponseWriter) {
	ErrorJSON(w, "body decoding error", http.StatusBadRequest)
}
