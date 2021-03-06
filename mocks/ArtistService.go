// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	app_errors "setmaker-api-go-rest/internal/utils/errors"

	domain "setmaker-api-go-rest/internal/domain"

	mock "github.com/stretchr/testify/mock"

	utils "setmaker-api-go-rest/internal/utils"

	uuid "github.com/google/uuid"
)

// ArtistService is an autogenerated mock type for the ArtistService type
type ArtistService struct {
	mock.Mock
}

// CreateArtist provides a mock function with given fields: _a0, _a1
func (_m *ArtistService) CreateArtist(_a0 context.Context, _a1 *domain.Artist) *app_errors.AppError {
	ret := _m.Called(_a0, _a1)

	var r0 *app_errors.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Artist) *app_errors.AppError); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app_errors.AppError)
		}
	}

	return r0
}

// DeleteArtist provides a mock function with given fields: _a0, _a1
func (_m *ArtistService) DeleteArtist(_a0 context.Context, _a1 uuid.UUID) (*domain.Artist, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Artist
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Artist); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Artist)
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

// GetArtist provides a mock function with given fields: _a0, _a1
func (_m *ArtistService) GetArtist(_a0 context.Context, _a1 uuid.UUID) (*domain.Artist, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 *domain.Artist
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *domain.Artist); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Artist)
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

// GetArtists provides a mock function with given fields: _a0, _a1
func (_m *ArtistService) GetArtists(_a0 context.Context, _a1 *utils.QuerySort) ([]*domain.Artist, *app_errors.AppError) {
	ret := _m.Called(_a0, _a1)

	var r0 []*domain.Artist
	if rf, ok := ret.Get(0).(func(context.Context, *utils.QuerySort) []*domain.Artist); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Artist)
		}
	}

	var r1 *app_errors.AppError
	if rf, ok := ret.Get(1).(func(context.Context, *utils.QuerySort) *app_errors.AppError); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*app_errors.AppError)
		}
	}

	return r0, r1
}

// UpdateArtist provides a mock function with given fields: _a0, _a1, _a2
func (_m *ArtistService) UpdateArtist(_a0 context.Context, _a1 *domain.Artist, _a2 uuid.UUID) *app_errors.AppError {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *app_errors.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Artist, uuid.UUID) *app_errors.AppError); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app_errors.AppError)
		}
	}

	return r0
}
