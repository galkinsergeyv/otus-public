// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package handle_exceptions is a generated GoMock package.
package handle_exceptions

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockILogger is a mock of ILogger interface.
type MockILogger struct {
	ctrl     *gomock.Controller
	recorder *MockILoggerMockRecorder
}

// MockILoggerMockRecorder is the mock recorder for MockILogger.
type MockILoggerMockRecorder struct {
	mock *MockILogger
}

// NewMockILogger creates a new mock instance.
func NewMockILogger(ctrl *gomock.Controller) *MockILogger {
	mock := &MockILogger{ctrl: ctrl}
	mock.recorder = &MockILoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILogger) EXPECT() *MockILoggerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockILogger) Log(msg string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Log", msg)
}

// Log indicates an expected call of Log.
func (mr *MockILoggerMockRecorder) Log(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockILogger)(nil).Log), msg)
}

// MockICommand is a mock of ICommand interface.
type MockICommand struct {
	ctrl     *gomock.Controller
	recorder *MockICommandMockRecorder
}

// MockICommandMockRecorder is the mock recorder for MockICommand.
type MockICommandMockRecorder struct {
	mock *MockICommand
}

// NewMockICommand creates a new mock instance.
func NewMockICommand(ctrl *gomock.Controller) *MockICommand {
	mock := &MockICommand{ctrl: ctrl}
	mock.recorder = &MockICommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommand) EXPECT() *MockICommandMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockICommand) Execute() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute")
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockICommandMockRecorder) Execute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockICommand)(nil).Execute))
}

// GetType mocks base method.
func (m *MockICommand) GetType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockICommandMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockICommand)(nil).GetType))
}

// MockIExceptionHandler is a mock of IExceptionHandler interface.
type MockIExceptionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIExceptionHandlerMockRecorder
}

// MockIExceptionHandlerMockRecorder is the mock recorder for MockIExceptionHandler.
type MockIExceptionHandlerMockRecorder struct {
	mock *MockIExceptionHandler
}

// NewMockIExceptionHandler creates a new mock instance.
func NewMockIExceptionHandler(ctrl *gomock.Controller) *MockIExceptionHandler {
	mock := &MockIExceptionHandler{ctrl: ctrl}
	mock.recorder = &MockIExceptionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIExceptionHandler) EXPECT() *MockIExceptionHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockIExceptionHandler) Handle(cmd ICommand, err error, queue chan ICommand) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", cmd, err, queue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIExceptionHandlerMockRecorder) Handle(cmd, err, queue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIExceptionHandler)(nil).Handle), cmd, err, queue)
}
