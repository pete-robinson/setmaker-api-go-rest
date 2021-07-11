package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * Templated json response
 */
func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/**
 * create an error payload
 */
func CreateErrorPayload(e error) map[string]string {
	fmt.Println("HERE")
	return map[string]string{
		"error": e.Error(),
	}
}
