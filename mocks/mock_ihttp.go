// Code generated by MockGen. DO NOT EDIT.
// Source: ihttp.go

// Package mocks is a generated GoMock package.
package mocks

import (
	resources "github.com/astronautsid/astro-ims-be/resources"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIHTTP is a mock of IHTTP interface.
type MockIHTTP struct {
	ctrl     *gomock.Controller
	recorder *MockIHTTPMockRecorder
}

// MockIHTTPMockRecorder is the mock recorder for MockIHTTP.
type MockIHTTPMockRecorder struct {
	mock *MockIHTTP
}

// NewMockIHTTP creates a new mock instance.
func NewMockIHTTP(ctrl *gomock.Controller) *MockIHTTP {
	mock := &MockIHTTP{ctrl: ctrl}
	mock.recorder = &MockIHTTPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHTTP) EXPECT() *MockIHTTPMockRecorder {
	return m.recorder
}

// CallService mocks base method.
func (m *MockIHTTP) CallService(method, url string, requestBody []byte) (string, *resources.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallService", method, url, requestBody)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*resources.ApplicationError)
	return ret0, ret1
}

// CallService indicates an expected call of CallService.
func (mr *MockIHTTPMockRecorder) CallService(method, url, requestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallService", reflect.TypeOf((*MockIHTTP)(nil).CallService), method, url, requestBody)
}

// CallServiceByte mocks base method.
func (m *MockIHTTP) CallServiceByte(method, url string, requestBody []byte) ([]byte, *resources.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallServiceByte", method, url, requestBody)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*resources.ApplicationError)
	return ret0, ret1
}

// CallServiceByte indicates an expected call of CallServiceByte.
func (mr *MockIHTTPMockRecorder) CallServiceByte(method, url, requestBody interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallServiceByte", reflect.TypeOf((*MockIHTTP)(nil).CallServiceByte), method, url, requestBody)
}
