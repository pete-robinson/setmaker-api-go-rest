// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	app_errors "setmaker-api-go-rest/internal/utils/errors"

	domain "setmaker-api-go-rest/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// SongService is an autogenerated mock type for the SongService type
type SongService struct {
	mock.Mock
}

// CreateSong provides a mock function with given fields: _a0, _a1
func (_m *SongService) CreateSong(_a0 context.Context, _a1 *domain.Song) *app_errors.AppError {
	ret := _m.Called(_a0, _a1)

	var r0 *app_errors.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Song) *app_errors.AppError); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app_errors.AppError)
		}
	}

	return r0
}

// DeleteSong provides a mock function with given fields: _a0, _a1
func (_m *SongService) DeleteSong(_a0 context.Context, _a1 uuid.UUID) (*domain.Song, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Song
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Song); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Song)
		}
	}

	var r1 *app_errors.AppError
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) *app_errors.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*app_errors.AppError)
		}
	}

	return r0, r1
}

// GetSong provides a mock function with given fields: _a0, _a1
func (_m *SongService) GetSong(_a0 context.Context, _a1 uuid.UUID) (*domain.Song, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Song
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Song); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Song)
		}
	}

	var r1 *app_errors.AppError
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) *app_errors.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*app_errors.AppError)
		}
	}

	return r0, r1
}

// GetSongsByArtistId provides a mock function with given fields: _a0, _a1
func (_m *SongService) GetSongsByArtistId(_a0 context.Context, _a1 uuid.UUID) ([]*domain.Song, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Song
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*domain.Song); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Song)
		}
	}

	var r1 *app_errors.AppError
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) *app_errors.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*app_errors.AppError)
		}
	}

	return r0, r1
}

// UpdateSong provides a mock function with given fields: _a0, _a1, _a2
func (_m *SongService) UpdateSong(_a0 context.Context, _a1 *domain.Song, _a2 uuid.UUID) *app_errors.AppError {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *app_errors.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Song, uuid.UUID) *app_errors.AppError); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app_errors.AppError)
		}
	}

	return r0
}
