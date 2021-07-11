package app_errors

import (
	"encoding/json"
	"errors"
)

type ErrorPayload struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

type BaseError struct {
	Code    int          `json:"status"`
	Payload ErrorPayload `json:"payload"`
}

func (a *BaseError) GetPayload() ErrorPayload {
	return a.Payload
}

func (a *BaseError) GetCode() int {
	return a.Code
}

func (a *BaseError) GetMessage() (interface{}, error) {
	if (ErrorPayload{}) == a.Payload {
		return nil, errors.New("Error payload not found")
	}

	return a.Payload.Message, nil
}

/**
 * @todo deal with how Message is handled
 * json won't unmarshall a type interface{}...
 */
func (a *BaseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    int         `json:"status"`
		Error   string      `json:"error"`
		Message interface{} `json:"message"`
	}{
		a.Code,
		a.Payload.Error,
		a.Payload.Message,
	})
}
