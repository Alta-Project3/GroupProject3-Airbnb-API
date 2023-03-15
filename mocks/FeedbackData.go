// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	feedback "groupproject3-airbnb-api/features/feedback"

	mock "github.com/stretchr/testify/mock"
)

// FeedbackData is an autogenerated mock type for the FeedbackData type
type FeedbackData struct {
	mock.Mock
}

// Create provides a mock function with given fields: userID, roomID, newFeedback
func (_m *FeedbackData) Create(userID uint, roomID uint, newFeedback feedback.Core) (feedback.Core, error) {
	ret := _m.Called(userID, roomID, newFeedback)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(uint, uint, feedback.Core) feedback.Core); ok {
		r0 = rf(userID, roomID, newFeedback)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, feedback.Core) error); ok {
		r1 = rf(userID, roomID, newFeedback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: userID, feedbackID
func (_m *FeedbackData) GetByID(userID uint, feedbackID uint) (feedback.Core, error) {
	ret := _m.Called(userID, feedbackID)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(uint, uint) feedback.Core); ok {
		r0 = rf(userID, feedbackID)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, feedbackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserFeedback provides a mock function with given fields: userID
func (_m *FeedbackData) GetUserFeedback(userID uint) ([]feedback.Core, error) {
	ret := _m.Called(userID)

	var r0 []feedback.Core
	if rf, ok := ret.Get(0).(func(uint) []feedback.Core); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feedback.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, feedBackID
func (_m *FeedbackData) Update(userID uint, feedBackID uint) (feedback.Core, error) {
	ret := _m.Called(userID, feedBackID)

	var r0 feedback.Core
	if rf, ok := ret.Get(0).(func(uint, uint) feedback.Core); ok {
		r0 = rf(userID, feedBackID)
	} else {
		r0 = ret.Get(0).(feedback.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userID, feedBackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeedbackData interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeedbackData creates a new instance of FeedbackData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeedbackData(t mockConstructorTestingTNewFeedbackData) *FeedbackData {
	mock := &FeedbackData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
