// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "setmaker-api-go-rest/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// SongsRepository is an autogenerated mock type for the SongsRepository type
type SongsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *SongsRepository) Create(_a0 context.Context, _a1 *domain.Song) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Song) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SongsRepository) Delete(_a0 context.Context, _a1 *domain.Song) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Song) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindSongsByArtistId provides a mock function with given fields: _a0, _a1
func (_m *SongsRepository) FindSongsByArtistId(_a0 context.Context, _a1 uuid.UUID) ([]*domain.Song, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Song
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*domain.Song); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Song)
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

// GetById provides a mock function with given fields: _a0, _a1
func (_m *SongsRepository) GetById(_a0 context.Context, _a1 uuid.UUID) (*domain.Song, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Song
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Song); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Song)
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
func (_m *SongsRepository) Update(_a0 context.Context, _a1 *domain.Song) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Song) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}