// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	feedback "groupproject3-airbnb-api/features/feedback"

	mock "github.com/stretchr/testify/mock"
)

// FeedbackService is an autogenerated mock type for the FeedbackService type
type FeedbackService struct {
	mock.Mock
}

// Create provides a mock function with given fields: token, roomID, newFeedback
func (_m *FeedbackService) Create(token interface{}, roomID uint, newFeedback feedback.Core) (feedback.Core, error) {
	ret := _m.Called(token, roomID, newFeedback)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, feedback.Core) feedback.Core); ok {
		r0 = rf(token, roomID, newFeedback)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, feedback.Core) error); ok {
		r1 = rf(token, roomID, newFeedback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: token, feedbackID
func (_m *FeedbackService) GetByID(token interface{}, feedbackID uint) (feedback.Core, error) {
	ret := _m.Called(token, feedbackID)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint) feedback.Core); ok {
		r0 = rf(token, feedbackID)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint) error); ok {
		r1 = rf(token, feedbackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFeedback provides a mock function with given fields: token
func (_m *FeedbackService) GetUserFeedback(token interface{}) ([]feedback.Core, error) {
	ret := _m.Called(token)

	var r0 []feedback.Core
	if rf, ok := ret.Get(0).(func(interface{}) []feedback.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feedback.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: token, feedbackID, updatedFeedback
func (_m *FeedbackService) Update(token interface{}, feedbackID uint, updatedFeedback feedback.Core) (feedback.Core, error) {
	ret := _m.Called(token, feedbackID, updatedFeedback)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, feedback.Core) feedback.Core); ok {
		r0 = rf(token, feedbackID, updatedFeedback)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, feedback.Core) error); ok {
		r1 = rf(token, feedbackID, updatedFeedback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeedbackService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeedbackService creates a new instance of FeedbackService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeedbackService(t mockConstructorTestingTNewFeedbackService) *FeedbackService {
	mock := &FeedbackService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}