package handlers

import (
	"encoding/json"
	"net/http"
)

type JSONError struct {
	Err            string `json:"error,omitempty"`
	ErrDescription string `json:"err_description,omitempty"`
}

func JSONResponse(w http.ResponseWriter, data any, code int) {
	json.NewEncoder(w).Encode(data)
}

func JSONErrorResponse(w http.ResponseWriter, jsonErr JSONError, code int) {
	JSONResponse(w, jsonErr, code)
}

func JSONInternalErrorResponse(w http.ResponseWriter) {
	JSONErrorResponse(w, JSONError{
		Err: "internal server error",
	}, http.StatusInternalServerError)
}
