// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_register is a generated GoMock package.
package mock_register

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	register "github.com/solo-io/skv2/pkg/multicluster/register"
	v1 "k8s.io/api/core/v1"
	rest "k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterRegistrant is a mock of ClusterRegistrant interface
type MockClusterRegistrant struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistrantMockRecorder
}

// MockClusterRegistrantMockRecorder is the mock recorder for MockClusterRegistrant
type MockClusterRegistrantMockRecorder struct {
	mock *MockClusterRegistrant
}

// NewMockClusterRegistrant creates a new mock instance
func NewMockClusterRegistrant(ctrl *gomock.Controller) *MockClusterRegistrant {
	mock := &MockClusterRegistrant{ctrl: ctrl}
	mock.recorder = &MockClusterRegistrantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterRegistrant) EXPECT() *MockClusterRegistrantMockRecorder {
	return m.recorder
}

// EnsureRemoteServiceAccount mocks base method
func (m *MockClusterRegistrant) EnsureRemoteServiceAccount(ctx context.Context, remoteClientCfg clientcmd.ClientConfig, opts register.Options) (*v1.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureRemoteServiceAccount", ctx, remoteClientCfg, opts)
	ret0, _ := ret[0].(*v1.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureRemoteServiceAccount indicates an expected call of EnsureRemoteServiceAccount
func (mr *MockClusterRegistrantMockRecorder) EnsureRemoteServiceAccount(ctx, remoteClientCfg, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureRemoteServiceAccount", reflect.TypeOf((*MockClusterRegistrant)(nil).EnsureRemoteServiceAccount), ctx, remoteClientCfg, opts)
}

// CreateRemoteAccessToken mocks base method
func (m *MockClusterRegistrant) CreateRemoteAccessToken(ctx context.Context, remoteClientCfg clientcmd.ClientConfig, sa client.ObjectKey, opts register.RbacOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRemoteAccessToken", ctx, remoteClientCfg, sa, opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRemoteAccessToken indicates an expected call of CreateRemoteAccessToken
func (mr *MockClusterRegistrantMockRecorder) CreateRemoteAccessToken(ctx, remoteClientCfg, sa, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRemoteAccessToken", reflect.TypeOf((*MockClusterRegistrant)(nil).CreateRemoteAccessToken), ctx, remoteClientCfg, sa, opts)
}

// RegisterClusterWithToken mocks base method
func (m *MockClusterRegistrant) RegisterClusterWithToken(ctx context.Context, masterClusterCfg *rest.Config, remoteClientCfg clientcmd.ClientConfig, token string, opts register.Options) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClusterWithToken", ctx, masterClusterCfg, remoteClientCfg, token, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterClusterWithToken indicates an expected call of RegisterClusterWithToken
func (mr *MockClusterRegistrantMockRecorder) RegisterClusterWithToken(ctx, masterClusterCfg, remoteClientCfg, token, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClusterWithToken", reflect.TypeOf((*MockClusterRegistrant)(nil).RegisterClusterWithToken), ctx, masterClusterCfg, remoteClientCfg, token, opts)
}
