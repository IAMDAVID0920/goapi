package api

import (
	"encoding/json"
	"net/http"
)

// import (
// 	"encoding/json"
// 	"net/http"
// )

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Success code return, usually 200
	Code int
	// Your account balance
	Balance int64
}

type Error struct {
	// Error code
	Code int
	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// We want to create multiple different error handlers
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An unexpected error occurred...", http.StatusBadRequest) // 400 error bad request
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An internal error occurred...", http.StatusInternalServerError) // 500 internal error code
	}
)
