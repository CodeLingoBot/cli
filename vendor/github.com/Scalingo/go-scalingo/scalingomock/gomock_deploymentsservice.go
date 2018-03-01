// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: DeploymentsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	websocket "golang.org/x/net/websocket"
	http "net/http"
	reflect "reflect"
)

// MockDeploymentsService is a mock of DeploymentsService interface
type MockDeploymentsService struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentsServiceMockRecorder
}

// MockDeploymentsServiceMockRecorder is the mock recorder for MockDeploymentsService
type MockDeploymentsServiceMockRecorder struct {
	mock *MockDeploymentsService
}

// NewMockDeploymentsService creates a new mock instance
func NewMockDeploymentsService(ctrl *gomock.Controller) *MockDeploymentsService {
	mock := &MockDeploymentsService{ctrl: ctrl}
	mock.recorder = &MockDeploymentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeploymentsService) EXPECT() *MockDeploymentsServiceMockRecorder {
	return m.recorder
}

// Deployment mocks base method
func (m *MockDeploymentsService) Deployment(arg0, arg1 string) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "Deployment", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deployment indicates an expected call of Deployment
func (mr *MockDeploymentsServiceMockRecorder) Deployment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockDeploymentsService)(nil).Deployment), arg0, arg1)
}

// DeploymentList mocks base method
func (m *MockDeploymentsService) DeploymentList(arg0 string) ([]*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentList indicates an expected call of DeploymentList
func (mr *MockDeploymentsServiceMockRecorder) DeploymentList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentList", reflect.TypeOf((*MockDeploymentsService)(nil).DeploymentList), arg0)
}

// DeploymentLogs mocks base method
func (m *MockDeploymentsService) DeploymentLogs(arg0 string) (*http.Response, error) {
	ret := m.ctrl.Call(m, "DeploymentLogs", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentLogs indicates an expected call of DeploymentLogs
func (mr *MockDeploymentsServiceMockRecorder) DeploymentLogs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentLogs", reflect.TypeOf((*MockDeploymentsService)(nil).DeploymentLogs), arg0)
}

// DeploymentStream mocks base method
func (m *MockDeploymentsService) DeploymentStream(arg0 string) (*websocket.Conn, error) {
	ret := m.ctrl.Call(m, "DeploymentStream", arg0)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentStream indicates an expected call of DeploymentStream
func (mr *MockDeploymentsServiceMockRecorder) DeploymentStream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentStream", reflect.TypeOf((*MockDeploymentsService)(nil).DeploymentStream), arg0)
}

// DeploymentsCreate mocks base method
func (m *MockDeploymentsService) DeploymentsCreate(arg0 string, arg1 *go_scalingo.DeploymentsCreateParams) (*go_scalingo.Deployment, error) {
	ret := m.ctrl.Call(m, "DeploymentsCreate", arg0, arg1)
	ret0, _ := ret[0].(*go_scalingo.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentsCreate indicates an expected call of DeploymentsCreate
func (mr *MockDeploymentsServiceMockRecorder) DeploymentsCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentsCreate", reflect.TypeOf((*MockDeploymentsService)(nil).DeploymentsCreate), arg0, arg1)
}