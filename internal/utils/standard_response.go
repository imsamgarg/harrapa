package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(&data)

}

func SendServerSideErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(&DefaultServerErrorResponse)
}

func SendBadRequestResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(DefaultInvalidDataResponse)
}

var (
	DefaultServerErrorResponse     = NewErrorResponse("Server side error! Please try again later.")
	DefaultInvalidDataResponse     = NewErrorResponse("Invalid Details")
	DefaultUnauthenticatedResponse = NewErrorResponse("You are not authenticated")
	DefaultUnauthorizedResponse    = NewErrorResponse("You are not authorized")
)

func NewErrorResponse(msg string) map[string]string {
	return map[string]string{
		"error": msg,
	}
}
func NewSuccessResponse(msg string) map[string]string {
	return map[string]string{
		"msg": msg,
	}
}
