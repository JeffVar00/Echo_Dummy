// Code generated by mockery v2.16.0. DO NOT EDIT.

package service

import (
	models "Echo_Dummy/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// ShowPlayers provides a mock function with given fields: ctx, filename
func (_m *MockService) ShowPlayers(ctx context.Context, filename string) ([]models.Player, error) {
	ret := _m.Called(ctx, filename)

	var r0 []models.Player
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Player); ok {
		r0 = rf(ctx, filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Player)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
