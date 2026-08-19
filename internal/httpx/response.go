package httpx

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       any       `json:"data"`
	Error      *APIError `json:"error,omitempty"`
}

type APIError struct {
	Code    int    `json:"code"` // custom error code.
	Message string `json:"message"`
}

func OK(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	body := APIResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	}
	_ = json.NewEncoder(w).Encode(body)
}

const (
	CodeRefreshTokenInvalid = 11002
	CodePlayerNotFound      = 12002
)

var errorHTTPStatus = map[int]int{
	CodeRefreshTokenInvalid: http.StatusUnauthorized, // 401
	CodePlayerNotFound:      http.StatusNotFound,     // 404
}

func Fail(w http.ResponseWriter, code int, message string) {
	status, ok := errorHTTPStatus[code]
	if !ok {
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	apiErr := &APIError{Code: code, Message: message}
	body := APIResponse{
		StatusCode: status,
		Message:    "",
		Data:       nil,
		Error:      apiErr,
	}
	_ = json.NewEncoder(w).Encode(body)
}
