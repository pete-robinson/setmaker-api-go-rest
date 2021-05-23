package errors

import (
	"errors"
	"net/http"
)

type AppError interface {
	GetPayload() ErrorPayload
	GetCode() int
	GetMessage() (string, error)
}

type ErrorPayload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"-"`
}

type baseError struct {
	Code    int          `json:"code"`
	Payload ErrorPayload `json:"payload"`
}

type NotFound struct{ baseError }
type BadRequest struct{ baseError }
type Unauthorized struct{ baseError }
type Forbidden struct{ baseError }
type InternalServerError struct{ baseError }

func (a *baseError) GetPayload() ErrorPayload {
	return a.Payload
}

func (a *baseError) GetCode() int {
	return a.Code
}

func (a *baseError) GetMessage() (string, error) {
	if (ErrorPayload{}) == a.Payload {
		return "", errors.New("Error payload not found")
	}

	return a.Payload.Message, nil
}

func NewBadRequest(message string, data interface{}) *BadRequest {
	return &BadRequest{
		baseError{
			Code: http.StatusBadRequest,
			Payload: ErrorPayload{
				Message: message,
				Data:    data,
			},
		},
	}
}

func NewUnauthorized(message string, data interface{}) *Unauthorized {
	return &Unauthorized{
		baseError{
			Code: http.StatusUnauthorized,
			Payload: ErrorPayload{
				Message: message,
				Data:    data,
			},
		},
	}
}

func NewForbidden(message string, data interface{}) *Forbidden {
	return &Forbidden{
		baseError{
			Code: http.StatusForbidden,
			Payload: ErrorPayload{
				Message: message,
				Data:    data,
			},
		},
	}
}

func NewNotFound(message string, data interface{}) *NotFound {
	return &NotFound{
		baseError{
			Code: http.StatusNotFound,
			Payload: ErrorPayload{
				Message: message,
				Data:    data,
			},
		},
	}
}

func NewInternalServerError(message string, data interface{}) *InternalServerError {
	return &InternalServerError{
		baseError{
			Code: http.StatusInternalServerError,
			Payload: ErrorPayload{
				Message: message,
				Data:    data,
			},
		},
	}
}
