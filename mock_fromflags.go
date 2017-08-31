// Code generated by MockGen. DO NOT EDIT.
// Source: fromflags.go

// Package codec is a generated GoMock package.
package codec

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFromFlags is a mock of FromFlags interface
type MockFromFlags struct {
	ctrl     *gomock.Controller
	recorder *MockFromFlagsMockRecorder
}

// MockFromFlagsMockRecorder is the mock recorder for MockFromFlags
type MockFromFlagsMockRecorder struct {
	mock *MockFromFlags
}

// NewMockFromFlags creates a new mock instance
func NewMockFromFlags(ctrl *gomock.Controller) *MockFromFlags {
	mock := &MockFromFlags{ctrl: ctrl}
	mock.recorder = &MockFromFlagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFromFlags) EXPECT() *MockFromFlagsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockFromFlags) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockFromFlagsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockFromFlags)(nil).Validate))
}

// Make mocks base method
func (m *MockFromFlags) Make() Codec {
	ret := m.ctrl.Call(m, "Make")
	ret0, _ := ret[0].(Codec)
	return ret0
}

// Make indicates an expected call of Make
func (mr *MockFromFlagsMockRecorder) Make() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockFromFlags)(nil).Make))
}

// Type mocks base method
func (m *MockFromFlags) Type() string {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockFromFlagsMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockFromFlags)(nil).Type))
}
