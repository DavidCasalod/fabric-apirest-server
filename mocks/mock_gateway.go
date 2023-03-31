// Code generated by MockGen. DO NOT EDIT.
// Source: fabric/web (interfaces: GatewayInt)

// Package mocks is a generated GoMock package.
package fabric

import (
	fabric "fabric/web"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGatewayInt is a mock of GatewayInt interface.
type MockGatewayInt struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayIntMockRecorder
}

// MockGatewayIntMockRecorder is the mock recorder for MockGatewayInt.
type MockGatewayIntMockRecorder struct {
	mock *MockGatewayInt
}

// NewMockGatewayInt creates a new mock instance.
func NewMockGatewayInt(ctrl *gomock.Controller) *MockGatewayInt {
	mock := &MockGatewayInt{ctrl: ctrl}
	mock.recorder = &MockGatewayIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayInt) EXPECT() *MockGatewayIntMockRecorder {
	return m.recorder
}

// GetNetwork mocks base method.
func (m *MockGatewayInt) GetNetwork(arg0 string) fabric.NetworkInt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", arg0)
	ret0, _ := ret[0].(fabric.NetworkInt)
	return ret0
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockGatewayIntMockRecorder) GetNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockGatewayInt)(nil).GetNetwork), arg0)
}
