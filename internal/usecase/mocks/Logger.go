// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: msg, fields
func (_m *Logger) Debug(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg, fields
func (_m *Logger) Error(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: msg, fields
func (_m *Logger) Fatal(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: msg, fields
func (_m *Logger) Info(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: msg, fields
func (_m *Logger) Warn(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
