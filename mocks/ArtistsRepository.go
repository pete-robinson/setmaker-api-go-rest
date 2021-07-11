// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "setmaker-api-go-rest/internal/domain"

	mock "github.com/stretchr/testify/mock"

	utils "setmaker-api-go-rest/internal/utils"

	uuid "github.com/google/uuid"
)

// ArtistsRepository is an autogenerated mock type for the ArtistsRepository type
type ArtistsRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) Count(_a0 context.Context, _a1 ...utils.FieldSearch) (int64, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, ...utils.FieldSearch) int64); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...utils.FieldSearch) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) Create(_a0 context.Context, _a1 *domain.Artist) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Artist) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) Delete(_a0 context.Context, _a1 *domain.Artist) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Artist) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) Find(_a0 context.Context, _a1 *utils.QuerySort) ([]*domain.Artist, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Artist
	if rf, ok := ret.Get(0).(func(context.Context, *utils.QuerySort) []*domain.Artist); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Artist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *utils.QuerySort) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) GetById(_a0 context.Context, _a1 uuid.UUID) (*domain.Artist, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Artist
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Artist); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Artist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ArtistsRepository) Update(_a0 context.Context, _a1 *domain.Artist) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Artist) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
