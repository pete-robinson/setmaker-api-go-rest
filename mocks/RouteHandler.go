// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RouteHandler is an autogenerated mock type for the RouteHandler type
type RouteHandler struct {
	mock.Mock
}

// HandleRoutes provides a mock function with given fields: _a0, _a1
func (_m *RouteHandler) HandleRoutes(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}
