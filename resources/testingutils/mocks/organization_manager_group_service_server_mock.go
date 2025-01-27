// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1 (interfaces: GroupServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	organizationmanager "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1"
	reflect "reflect"
)

// MockGroupServiceServer is a mock of GroupServiceServer interface
type MockGroupServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceServerMockRecorder
}

// MockGroupServiceServerMockRecorder is the mock recorder for MockGroupServiceServer
type MockGroupServiceServerMockRecorder struct {
	mock *MockGroupServiceServer
}

// NewMockGroupServiceServer creates a new mock instance
func NewMockGroupServiceServer(ctrl *gomock.Controller) *MockGroupServiceServer {
	mock := &MockGroupServiceServer{ctrl: ctrl}
	mock.recorder = &MockGroupServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupServiceServer) EXPECT() *MockGroupServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockGroupServiceServer) Create(arg0 context.Context, arg1 *organizationmanager.CreateGroupRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGroupServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGroupServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockGroupServiceServer) Delete(arg0 context.Context, arg1 *organizationmanager.DeleteGroupRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockGroupServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockGroupServiceServer) Get(arg0 context.Context, arg1 *organizationmanager.GetGroupRequest) (*organizationmanager.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*organizationmanager.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGroupServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGroupServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockGroupServiceServer) List(arg0 context.Context, arg1 *organizationmanager.ListGroupsRequest) (*organizationmanager.ListGroupsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*organizationmanager.ListGroupsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGroupServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGroupServiceServer)(nil).List), arg0, arg1)
}

// ListAccessBindings mocks base method
func (m *MockGroupServiceServer) ListAccessBindings(arg0 context.Context, arg1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*access.ListAccessBindingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessBindings indicates an expected call of ListAccessBindings
func (mr *MockGroupServiceServerMockRecorder) ListAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessBindings", reflect.TypeOf((*MockGroupServiceServer)(nil).ListAccessBindings), arg0, arg1)
}

// ListMembers mocks base method
func (m *MockGroupServiceServer) ListMembers(arg0 context.Context, arg1 *organizationmanager.ListGroupMembersRequest) (*organizationmanager.ListGroupMembersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", arg0, arg1)
	ret0, _ := ret[0].(*organizationmanager.ListGroupMembersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers
func (mr *MockGroupServiceServerMockRecorder) ListMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockGroupServiceServer)(nil).ListMembers), arg0, arg1)
}

// ListOperations mocks base method
func (m *MockGroupServiceServer) ListOperations(arg0 context.Context, arg1 *organizationmanager.ListGroupOperationsRequest) (*organizationmanager.ListGroupOperationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1)
	ret0, _ := ret[0].(*organizationmanager.ListGroupOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations
func (mr *MockGroupServiceServerMockRecorder) ListOperations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockGroupServiceServer)(nil).ListOperations), arg0, arg1)
}

// SetAccessBindings mocks base method
func (m *MockGroupServiceServer) SetAccessBindings(arg0 context.Context, arg1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccessBindings indicates an expected call of SetAccessBindings
func (mr *MockGroupServiceServerMockRecorder) SetAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessBindings", reflect.TypeOf((*MockGroupServiceServer)(nil).SetAccessBindings), arg0, arg1)
}

// Update mocks base method
func (m *MockGroupServiceServer) Update(arg0 context.Context, arg1 *organizationmanager.UpdateGroupRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockGroupServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupServiceServer)(nil).Update), arg0, arg1)
}

// UpdateAccessBindings mocks base method
func (m *MockGroupServiceServer) UpdateAccessBindings(arg0 context.Context, arg1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccessBindings indicates an expected call of UpdateAccessBindings
func (mr *MockGroupServiceServerMockRecorder) UpdateAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessBindings", reflect.TypeOf((*MockGroupServiceServer)(nil).UpdateAccessBindings), arg0, arg1)
}

// UpdateMembers mocks base method
func (m *MockGroupServiceServer) UpdateMembers(arg0 context.Context, arg1 *organizationmanager.UpdateGroupMembersRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMembers", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMembers indicates an expected call of UpdateMembers
func (mr *MockGroupServiceServerMockRecorder) UpdateMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMembers", reflect.TypeOf((*MockGroupServiceServer)(nil).UpdateMembers), arg0, arg1)
}
