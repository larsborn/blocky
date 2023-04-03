// Code generated by mockery v2.23.1. DO NOT EDIT.

package api

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockBlockingControl is an autogenerated mock type for the BlockingControl type
type MockBlockingControl struct {
	mock.Mock
}

type MockBlockingControl_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBlockingControl) EXPECT() *MockBlockingControl_Expecter {
	return &MockBlockingControl_Expecter{mock: &_m.Mock}
}

// BlockingStatus provides a mock function with given fields:
func (_m *MockBlockingControl) BlockingStatus() BlockingStatus {
	ret := _m.Called()

	var r0 BlockingStatus
	if rf, ok := ret.Get(0).(func() BlockingStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BlockingStatus)
	}

	return r0
}

// MockBlockingControl_BlockingStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockingStatus'
type MockBlockingControl_BlockingStatus_Call struct {
	*mock.Call
}

// BlockingStatus is a helper method to define mock.On call
func (_e *MockBlockingControl_Expecter) BlockingStatus() *MockBlockingControl_BlockingStatus_Call {
	return &MockBlockingControl_BlockingStatus_Call{Call: _e.mock.On("BlockingStatus")}
}

func (_c *MockBlockingControl_BlockingStatus_Call) Run(run func()) *MockBlockingControl_BlockingStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBlockingControl_BlockingStatus_Call) Return(_a0 BlockingStatus) *MockBlockingControl_BlockingStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlockingControl_BlockingStatus_Call) RunAndReturn(run func() BlockingStatus) *MockBlockingControl_BlockingStatus_Call {
	_c.Call.Return(run)
	return _c
}

// DisableBlocking provides a mock function with given fields: duration, disableGroups
func (_m *MockBlockingControl) DisableBlocking(duration time.Duration, disableGroups []string) error {
	ret := _m.Called(duration, disableGroups)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, []string) error); ok {
		r0 = rf(duration, disableGroups)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBlockingControl_DisableBlocking_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DisableBlocking'
type MockBlockingControl_DisableBlocking_Call struct {
	*mock.Call
}

// DisableBlocking is a helper method to define mock.On call
//   - duration time.Duration
//   - disableGroups []string
func (_e *MockBlockingControl_Expecter) DisableBlocking(duration interface{}, disableGroups interface{}) *MockBlockingControl_DisableBlocking_Call {
	return &MockBlockingControl_DisableBlocking_Call{Call: _e.mock.On("DisableBlocking", duration, disableGroups)}
}

func (_c *MockBlockingControl_DisableBlocking_Call) Run(run func(duration time.Duration, disableGroups []string)) *MockBlockingControl_DisableBlocking_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].([]string))
	})
	return _c
}

func (_c *MockBlockingControl_DisableBlocking_Call) Return(_a0 error) *MockBlockingControl_DisableBlocking_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlockingControl_DisableBlocking_Call) RunAndReturn(run func(time.Duration, []string) error) *MockBlockingControl_DisableBlocking_Call {
	_c.Call.Return(run)
	return _c
}

// EnableBlocking provides a mock function with given fields:
func (_m *MockBlockingControl) EnableBlocking() {
	_m.Called()
}

// MockBlockingControl_EnableBlocking_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnableBlocking'
type MockBlockingControl_EnableBlocking_Call struct {
	*mock.Call
}

// EnableBlocking is a helper method to define mock.On call
func (_e *MockBlockingControl_Expecter) EnableBlocking() *MockBlockingControl_EnableBlocking_Call {
	return &MockBlockingControl_EnableBlocking_Call{Call: _e.mock.On("EnableBlocking")}
}

func (_c *MockBlockingControl_EnableBlocking_Call) Run(run func()) *MockBlockingControl_EnableBlocking_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBlockingControl_EnableBlocking_Call) Return() *MockBlockingControl_EnableBlocking_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBlockingControl_EnableBlocking_Call) RunAndReturn(run func()) *MockBlockingControl_EnableBlocking_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockBlockingControl interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBlockingControl creates a new instance of MockBlockingControl. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBlockingControl(t mockConstructorTestingTNewMockBlockingControl) *MockBlockingControl {
	mock := &MockBlockingControl{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}