// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openfga/api/proto/openfga/v1 (interfaces: OpenFGAServiceClient)
//
// Generated by this command:
//
//	mockgen -package mocks -destination internal/mocks/mock_openfga_service_client.go github.com/openfga/api/proto/openfga/v1 OpenFGAServiceClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockOpenFGAServiceClient is a mock of OpenFGAServiceClient interface.
type MockOpenFGAServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenFGAServiceClientMockRecorder
}

// MockOpenFGAServiceClientMockRecorder is the mock recorder for MockOpenFGAServiceClient.
type MockOpenFGAServiceClientMockRecorder struct {
	mock *MockOpenFGAServiceClient
}

// NewMockOpenFGAServiceClient creates a new mock instance.
func NewMockOpenFGAServiceClient(ctrl *gomock.Controller) *MockOpenFGAServiceClient {
	mock := &MockOpenFGAServiceClient{ctrl: ctrl}
	mock.recorder = &MockOpenFGAServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenFGAServiceClient) EXPECT() *MockOpenFGAServiceClientMockRecorder {
	return m.recorder
}

// BatchCheck mocks base method.
func (m *MockOpenFGAServiceClient) BatchCheck(arg0 context.Context, arg1 *openfgav1.BatchCheckRequest, arg2 ...grpc.CallOption) (*openfgav1.BatchCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchCheck", varargs...)
	ret0, _ := ret[0].(*openfgav1.BatchCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchCheck indicates an expected call of BatchCheck.
func (mr *MockOpenFGAServiceClientMockRecorder) BatchCheck(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCheck", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).BatchCheck), varargs...)
}

// Check mocks base method.
func (m *MockOpenFGAServiceClient) Check(arg0 context.Context, arg1 *openfgav1.CheckRequest, arg2 ...grpc.CallOption) (*openfgav1.CheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Check", varargs...)
	ret0, _ := ret[0].(*openfgav1.CheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockOpenFGAServiceClientMockRecorder) Check(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).Check), varargs...)
}

// CreateStore mocks base method.
func (m *MockOpenFGAServiceClient) CreateStore(arg0 context.Context, arg1 *openfgav1.CreateStoreRequest, arg2 ...grpc.CallOption) (*openfgav1.CreateStoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStore", varargs...)
	ret0, _ := ret[0].(*openfgav1.CreateStoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStore indicates an expected call of CreateStore.
func (mr *MockOpenFGAServiceClientMockRecorder) CreateStore(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).CreateStore), varargs...)
}

// DeleteStore mocks base method.
func (m *MockOpenFGAServiceClient) DeleteStore(arg0 context.Context, arg1 *openfgav1.DeleteStoreRequest, arg2 ...grpc.CallOption) (*openfgav1.DeleteStoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStore", varargs...)
	ret0, _ := ret[0].(*openfgav1.DeleteStoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStore indicates an expected call of DeleteStore.
func (mr *MockOpenFGAServiceClientMockRecorder) DeleteStore(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStore", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).DeleteStore), varargs...)
}

// Expand mocks base method.
func (m *MockOpenFGAServiceClient) Expand(arg0 context.Context, arg1 *openfgav1.ExpandRequest, arg2 ...grpc.CallOption) (*openfgav1.ExpandResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Expand", varargs...)
	ret0, _ := ret[0].(*openfgav1.ExpandResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expand indicates an expected call of Expand.
func (mr *MockOpenFGAServiceClientMockRecorder) Expand(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expand", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).Expand), varargs...)
}

// GetStore mocks base method.
func (m *MockOpenFGAServiceClient) GetStore(arg0 context.Context, arg1 *openfgav1.GetStoreRequest, arg2 ...grpc.CallOption) (*openfgav1.GetStoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStore", varargs...)
	ret0, _ := ret[0].(*openfgav1.GetStoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockOpenFGAServiceClientMockRecorder) GetStore(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).GetStore), varargs...)
}

// ListObjects mocks base method.
func (m *MockOpenFGAServiceClient) ListObjects(arg0 context.Context, arg1 *openfgav1.ListObjectsRequest, arg2 ...grpc.CallOption) (*openfgav1.ListObjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjects", varargs...)
	ret0, _ := ret[0].(*openfgav1.ListObjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects.
func (mr *MockOpenFGAServiceClientMockRecorder) ListObjects(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ListObjects), varargs...)
}

// ListStores mocks base method.
func (m *MockOpenFGAServiceClient) ListStores(arg0 context.Context, arg1 *openfgav1.ListStoresRequest, arg2 ...grpc.CallOption) (*openfgav1.ListStoresResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStores", varargs...)
	ret0, _ := ret[0].(*openfgav1.ListStoresResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStores indicates an expected call of ListStores.
func (mr *MockOpenFGAServiceClientMockRecorder) ListStores(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStores", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ListStores), varargs...)
}

// ListUsers mocks base method.
func (m *MockOpenFGAServiceClient) ListUsers(arg0 context.Context, arg1 *openfgav1.ListUsersRequest, arg2 ...grpc.CallOption) (*openfgav1.ListUsersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*openfgav1.ListUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockOpenFGAServiceClientMockRecorder) ListUsers(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ListUsers), varargs...)
}

// Read mocks base method.
func (m *MockOpenFGAServiceClient) Read(arg0 context.Context, arg1 *openfgav1.ReadRequest, arg2 ...grpc.CallOption) (*openfgav1.ReadResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*openfgav1.ReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockOpenFGAServiceClientMockRecorder) Read(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).Read), varargs...)
}

// ReadAssertions mocks base method.
func (m *MockOpenFGAServiceClient) ReadAssertions(arg0 context.Context, arg1 *openfgav1.ReadAssertionsRequest, arg2 ...grpc.CallOption) (*openfgav1.ReadAssertionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadAssertions", varargs...)
	ret0, _ := ret[0].(*openfgav1.ReadAssertionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAssertions indicates an expected call of ReadAssertions.
func (mr *MockOpenFGAServiceClientMockRecorder) ReadAssertions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAssertions", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ReadAssertions), varargs...)
}

// ReadAuthorizationModel mocks base method.
func (m *MockOpenFGAServiceClient) ReadAuthorizationModel(arg0 context.Context, arg1 *openfgav1.ReadAuthorizationModelRequest, arg2 ...grpc.CallOption) (*openfgav1.ReadAuthorizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadAuthorizationModel", varargs...)
	ret0, _ := ret[0].(*openfgav1.ReadAuthorizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthorizationModel indicates an expected call of ReadAuthorizationModel.
func (mr *MockOpenFGAServiceClientMockRecorder) ReadAuthorizationModel(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthorizationModel", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ReadAuthorizationModel), varargs...)
}

// ReadAuthorizationModels mocks base method.
func (m *MockOpenFGAServiceClient) ReadAuthorizationModels(arg0 context.Context, arg1 *openfgav1.ReadAuthorizationModelsRequest, arg2 ...grpc.CallOption) (*openfgav1.ReadAuthorizationModelsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadAuthorizationModels", varargs...)
	ret0, _ := ret[0].(*openfgav1.ReadAuthorizationModelsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAuthorizationModels indicates an expected call of ReadAuthorizationModels.
func (mr *MockOpenFGAServiceClientMockRecorder) ReadAuthorizationModels(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAuthorizationModels", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ReadAuthorizationModels), varargs...)
}

// ReadChanges mocks base method.
func (m *MockOpenFGAServiceClient) ReadChanges(arg0 context.Context, arg1 *openfgav1.ReadChangesRequest, arg2 ...grpc.CallOption) (*openfgav1.ReadChangesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadChanges", varargs...)
	ret0, _ := ret[0].(*openfgav1.ReadChangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadChanges indicates an expected call of ReadChanges.
func (mr *MockOpenFGAServiceClientMockRecorder) ReadChanges(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadChanges", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).ReadChanges), varargs...)
}

// StreamedListObjects mocks base method.
func (m *MockOpenFGAServiceClient) StreamedListObjects(arg0 context.Context, arg1 *openfgav1.StreamedListObjectsRequest, arg2 ...grpc.CallOption) (openfgav1.OpenFGAService_StreamedListObjectsClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamedListObjects", varargs...)
	ret0, _ := ret[0].(openfgav1.OpenFGAService_StreamedListObjectsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamedListObjects indicates an expected call of StreamedListObjects.
func (mr *MockOpenFGAServiceClientMockRecorder) StreamedListObjects(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamedListObjects", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).StreamedListObjects), varargs...)
}

// UpdateStore mocks base method.
func (m *MockOpenFGAServiceClient) UpdateStore(arg0 context.Context, arg1 *openfgav1.UpdateStoreRequest, arg2 ...grpc.CallOption) (*openfgav1.UpdateStoreResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateStore", varargs...)
	ret0, _ := ret[0].(*openfgav1.UpdateStoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStore indicates an expected call of UpdateStore.
func (mr *MockOpenFGAServiceClientMockRecorder) UpdateStore(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStore", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).UpdateStore), varargs...)
}

// Write mocks base method.
func (m *MockOpenFGAServiceClient) Write(arg0 context.Context, arg1 *openfgav1.WriteRequest, arg2 ...grpc.CallOption) (*openfgav1.WriteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Write", varargs...)
	ret0, _ := ret[0].(*openfgav1.WriteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockOpenFGAServiceClientMockRecorder) Write(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).Write), varargs...)
}

// WriteAssertions mocks base method.
func (m *MockOpenFGAServiceClient) WriteAssertions(arg0 context.Context, arg1 *openfgav1.WriteAssertionsRequest, arg2 ...grpc.CallOption) (*openfgav1.WriteAssertionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteAssertions", varargs...)
	ret0, _ := ret[0].(*openfgav1.WriteAssertionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAssertions indicates an expected call of WriteAssertions.
func (mr *MockOpenFGAServiceClientMockRecorder) WriteAssertions(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAssertions", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).WriteAssertions), varargs...)
}

// WriteAuthorizationModel mocks base method.
func (m *MockOpenFGAServiceClient) WriteAuthorizationModel(arg0 context.Context, arg1 *openfgav1.WriteAuthorizationModelRequest, arg2 ...grpc.CallOption) (*openfgav1.WriteAuthorizationModelResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WriteAuthorizationModel", varargs...)
	ret0, _ := ret[0].(*openfgav1.WriteAuthorizationModelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAuthorizationModel indicates an expected call of WriteAuthorizationModel.
func (mr *MockOpenFGAServiceClientMockRecorder) WriteAuthorizationModel(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAuthorizationModel", reflect.TypeOf((*MockOpenFGAServiceClient)(nil).WriteAuthorizationModel), varargs...)
}
