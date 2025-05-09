// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxidecomputer/rancher-machine-driver-oxide (interfaces: oxideClient)
//
// Generated by this command:
//
//	mockgen -destination mock/oxide_client.go -package mock . oxideClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	oxide "github.com/oxidecomputer/oxide.go/oxide"
	gomock "go.uber.org/mock/gomock"
)

// MockoxideClient is a mock of oxideClient interface.
type MockoxideClient struct {
	ctrl     *gomock.Controller
	recorder *MockoxideClientMockRecorder
	isgomock struct{}
}

// MockoxideClientMockRecorder is the mock recorder for MockoxideClient.
type MockoxideClientMockRecorder struct {
	mock *MockoxideClient
}

// NewMockoxideClient creates a new mock instance.
func NewMockoxideClient(ctrl *gomock.Controller) *MockoxideClient {
	mock := &MockoxideClient{ctrl: ctrl}
	mock.recorder = &MockoxideClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockoxideClient) EXPECT() *MockoxideClientMockRecorder {
	return m.recorder
}

// CurrentUserSshKeyCreate mocks base method.
func (m *MockoxideClient) CurrentUserSshKeyCreate(ctx context.Context, params oxide.CurrentUserSshKeyCreateParams) (*oxide.SshKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentUserSshKeyCreate", ctx, params)
	ret0, _ := ret[0].(*oxide.SshKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentUserSshKeyCreate indicates an expected call of CurrentUserSshKeyCreate.
func (mr *MockoxideClientMockRecorder) CurrentUserSshKeyCreate(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentUserSshKeyCreate", reflect.TypeOf((*MockoxideClient)(nil).CurrentUserSshKeyCreate), ctx, params)
}

// CurrentUserSshKeyDelete mocks base method.
func (m *MockoxideClient) CurrentUserSshKeyDelete(ctx context.Context, params oxide.CurrentUserSshKeyDeleteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentUserSshKeyDelete", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CurrentUserSshKeyDelete indicates an expected call of CurrentUserSshKeyDelete.
func (mr *MockoxideClientMockRecorder) CurrentUserSshKeyDelete(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentUserSshKeyDelete", reflect.TypeOf((*MockoxideClient)(nil).CurrentUserSshKeyDelete), ctx, params)
}

// DiskDelete mocks base method.
func (m *MockoxideClient) DiskDelete(ctx context.Context, params oxide.DiskDeleteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskDelete", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiskDelete indicates an expected call of DiskDelete.
func (mr *MockoxideClientMockRecorder) DiskDelete(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskDelete", reflect.TypeOf((*MockoxideClient)(nil).DiskDelete), ctx, params)
}

// InstanceCreate mocks base method.
func (m *MockoxideClient) InstanceCreate(ctx context.Context, params oxide.InstanceCreateParams) (*oxide.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceCreate", ctx, params)
	ret0, _ := ret[0].(*oxide.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceCreate indicates an expected call of InstanceCreate.
func (mr *MockoxideClientMockRecorder) InstanceCreate(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceCreate", reflect.TypeOf((*MockoxideClient)(nil).InstanceCreate), ctx, params)
}

// InstanceDelete mocks base method.
func (m *MockoxideClient) InstanceDelete(ctx context.Context, params oxide.InstanceDeleteParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceDelete", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstanceDelete indicates an expected call of InstanceDelete.
func (mr *MockoxideClientMockRecorder) InstanceDelete(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceDelete", reflect.TypeOf((*MockoxideClient)(nil).InstanceDelete), ctx, params)
}

// InstanceDiskListAllPages mocks base method.
func (m *MockoxideClient) InstanceDiskListAllPages(ctx context.Context, params oxide.InstanceDiskListParams) ([]oxide.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceDiskListAllPages", ctx, params)
	ret0, _ := ret[0].([]oxide.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceDiskListAllPages indicates an expected call of InstanceDiskListAllPages.
func (mr *MockoxideClientMockRecorder) InstanceDiskListAllPages(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceDiskListAllPages", reflect.TypeOf((*MockoxideClient)(nil).InstanceDiskListAllPages), ctx, params)
}

// InstanceNetworkInterfaceListAllPages mocks base method.
func (m *MockoxideClient) InstanceNetworkInterfaceListAllPages(ctx context.Context, params oxide.InstanceNetworkInterfaceListParams) ([]oxide.InstanceNetworkInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceNetworkInterfaceListAllPages", ctx, params)
	ret0, _ := ret[0].([]oxide.InstanceNetworkInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceNetworkInterfaceListAllPages indicates an expected call of InstanceNetworkInterfaceListAllPages.
func (mr *MockoxideClientMockRecorder) InstanceNetworkInterfaceListAllPages(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceNetworkInterfaceListAllPages", reflect.TypeOf((*MockoxideClient)(nil).InstanceNetworkInterfaceListAllPages), ctx, params)
}

// InstanceReboot mocks base method.
func (m *MockoxideClient) InstanceReboot(ctx context.Context, params oxide.InstanceRebootParams) (*oxide.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceReboot", ctx, params)
	ret0, _ := ret[0].(*oxide.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceReboot indicates an expected call of InstanceReboot.
func (mr *MockoxideClientMockRecorder) InstanceReboot(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceReboot", reflect.TypeOf((*MockoxideClient)(nil).InstanceReboot), ctx, params)
}

// InstanceStart mocks base method.
func (m *MockoxideClient) InstanceStart(ctx context.Context, params oxide.InstanceStartParams) (*oxide.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceStart", ctx, params)
	ret0, _ := ret[0].(*oxide.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceStart indicates an expected call of InstanceStart.
func (mr *MockoxideClientMockRecorder) InstanceStart(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceStart", reflect.TypeOf((*MockoxideClient)(nil).InstanceStart), ctx, params)
}

// InstanceStop mocks base method.
func (m *MockoxideClient) InstanceStop(ctx context.Context, params oxide.InstanceStopParams) (*oxide.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceStop", ctx, params)
	ret0, _ := ret[0].(*oxide.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceStop indicates an expected call of InstanceStop.
func (mr *MockoxideClientMockRecorder) InstanceStop(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceStop", reflect.TypeOf((*MockoxideClient)(nil).InstanceStop), ctx, params)
}

// InstanceView mocks base method.
func (m *MockoxideClient) InstanceView(ctx context.Context, params oxide.InstanceViewParams) (*oxide.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceView", ctx, params)
	ret0, _ := ret[0].(*oxide.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceView indicates an expected call of InstanceView.
func (mr *MockoxideClientMockRecorder) InstanceView(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceView", reflect.TypeOf((*MockoxideClient)(nil).InstanceView), ctx, params)
}
