// Code generated by MockGen. DO NOT EDIT.
// Source: throttler.go
//
// Generated by this command:
//
//	mockgen -source throttler.go -destination ../mocks/mock_throttler.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockThrottler is a mock of Throttler interface.
type MockThrottler struct {
	ctrl     *gomock.Controller
	recorder *MockThrottlerMockRecorder
}

// MockThrottlerMockRecorder is the mock recorder for MockThrottler.
type MockThrottlerMockRecorder struct {
	mock *MockThrottler
}

// NewMockThrottler creates a new mock instance.
func NewMockThrottler(ctrl *gomock.Controller) *MockThrottler {
	mock := &MockThrottler{ctrl: ctrl}
	mock.recorder = &MockThrottlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockThrottler) EXPECT() *MockThrottlerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockThrottler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockThrottlerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockThrottler)(nil).Close))
}

// Throttle mocks base method.
func (m *MockThrottler) Throttle(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Throttle", arg0)
}

// Throttle indicates an expected call of Throttle.
func (mr *MockThrottlerMockRecorder) Throttle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Throttle", reflect.TypeOf((*MockThrottler)(nil).Throttle), arg0)
}
