// Code generated by MockGen. DO NOT EDIT.
// Source: client/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "node_topology_discovery/model"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// MakeRequest mocks base method
func (m *MockClient) MakeRequest(ipAddress, port string, request model.NodesDiscoveryRequest) (model.NodesDiscoveryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", ipAddress, port, request)
	ret0, _ := ret[0].(model.NodesDiscoveryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest
func (mr *MockClientMockRecorder) MakeRequest(ipAddress, port, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockClient)(nil).MakeRequest), ipAddress, port, request)
}
