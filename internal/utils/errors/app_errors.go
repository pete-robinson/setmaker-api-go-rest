package app_errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	ErrCode int `json:"-"`
	BaseError
}

type IAppError interface {
	GetPayload() ErrorPayload
	GetCode() int
	GetMessage() (interface{}, error)
}

// IOTA for supported error types
const (
	ERRBadRequest = iota
	ERRUnauthorized
	ERRForbidden
	ERRNotFound
	ERRInternalServerError
)

/**
 * error strints - mapped to IOTA sequence above
 * maps to a map of the error name and code
 */
var errorStrings = map[int]map[string]interface{}{
	ERRBadRequest:          {"name": "Bad Request", "code": http.StatusBadRequest},
	ERRUnauthorized:        {"name": "Unauthorized", "code": http.StatusUnauthorized},
	ERRForbidden:           {"name": "Forbidden", "code": http.StatusForbidden},
	ERRNotFound:            {"name": "Not Found", "code": http.StatusNotFound},
	ERRInternalServerError: {"name": "Internal Server Error", "code": http.StatusInternalServerError},
}

/**
 * Make a new AppError based on an error reference and Message
 */
func MakeError(errRef int, message interface{}) *AppError {
	// base error obj
	code, ok := errorStrings[errRef]["code"].(int)
	if !ok {
		// code could not be converted to an int... use something generic
		code = 400
	}

	// construct the base error
	base := BaseError{
		Code: code,
		Payload: ErrorPayload{
			Error:   fmt.Sprintf("%v", errorStrings[errRef]["name"]),
			Message: message,
		},
	}

	// augment with the error reference
	err := &AppError{
		errRef,
		base,
	}

	return err
}
