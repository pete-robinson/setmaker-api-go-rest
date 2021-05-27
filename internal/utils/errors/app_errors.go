package app_errors

import (
	"net/http"
)

type AppError struct {
	ErrCode int `json:"-"`
	BaseError
}

// IOTA for supported error types
const (
	ERRBadRequest = iota
	ERRUnauthorized
	ERRForbidden
	ERRNotFound
	ERRInternalServerError
)

// error strints - mapped to IOTA sequence above
var errorStrings = map[int]string{
	ERRBadRequest:          "Bad Request",
	ERRUnauthorized:        "Unauthorized",
	ERRForbidden:           "Forbidden",
	ERRNotFound:            "Not Found",
	ERRInternalServerError: "Internal Server Error",
}

// @todo - map Code back to http status codes
func MakeError(errRef int, message interface{}) *AppError {
	// base error obj
	base := BaseError{
		Code: http.StatusBadRequest,
		Payload: ErrorPayload{
			Error:   errorStrings[errRef],
			Message: message,
		},
	}

	err := &AppError{
		errRef,
		base,
	}

	return err
}
