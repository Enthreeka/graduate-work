// Code generated by MockGen. DO NOT EDIT.
// Source: go/rpc_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=go/rpc_grpc.pb.go -destination=mock/rpc_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	proto "github.com/Entreeka/proto-proxy/go"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockGatewayClient is a mock of GatewayClient interface.
type MockGatewayClient struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayClientMockRecorder
}

// MockGatewayClientMockRecorder is the mock recorder for MockGatewayClient.
type MockGatewayClientMockRecorder struct {
	mock *MockGatewayClient
}

// NewMockGatewayClient creates a new mock instance.
func NewMockGatewayClient(ctrl *gomock.Controller) *MockGatewayClient {
	mock := &MockGatewayClient{ctrl: ctrl}
	mock.recorder = &MockGatewayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayClient) EXPECT() *MockGatewayClientMockRecorder {
	return m.recorder
}

// BulkAPI mocks base method.
func (m *MockGatewayClient) BulkAPI(ctx context.Context, in *proto.BulkAPIRequest, opts ...grpc.CallOption) (*proto.BulkAPIResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BulkAPI", varargs...)
	ret0, _ := ret[0].(*proto.BulkAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkAPI indicates an expected call of BulkAPI.
func (mr *MockGatewayClientMockRecorder) BulkAPI(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkAPI", reflect.TypeOf((*MockGatewayClient)(nil).BulkAPI), varargs...)
}

// CreateNewMovie mocks base method.
func (m *MockGatewayClient) CreateNewMovie(ctx context.Context, in *proto.CreateNewMovieRequest, opts ...grpc.CallOption) (*proto.CreateNewMovieResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewMovie", varargs...)
	ret0, _ := ret[0].(*proto.CreateNewMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewMovie indicates an expected call of CreateNewMovie.
func (mr *MockGatewayClientMockRecorder) CreateNewMovie(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewMovie", reflect.TypeOf((*MockGatewayClient)(nil).CreateNewMovie), varargs...)
}

// DeleteMovie mocks base method.
func (m *MockGatewayClient) DeleteMovie(ctx context.Context, in *proto.DeleteMovieRequest, opts ...grpc.CallOption) (*proto.DeleteMovieResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMovie", varargs...)
	ret0, _ := ret[0].(*proto.DeleteMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockGatewayClientMockRecorder) DeleteMovie(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockGatewayClient)(nil).DeleteMovie), varargs...)
}

// GetAllMovie mocks base method.
func (m *MockGatewayClient) GetAllMovie(ctx context.Context, in *proto.GetAllMovieRequest, opts ...grpc.CallOption) (*proto.GetAllMovieResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllMovie", varargs...)
	ret0, _ := ret[0].(*proto.GetAllMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMovie indicates an expected call of GetAllMovie.
func (mr *MockGatewayClientMockRecorder) GetAllMovie(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMovie", reflect.TypeOf((*MockGatewayClient)(nil).GetAllMovie), varargs...)
}

// GetIndices mocks base method.
func (m *MockGatewayClient) GetIndices(ctx context.Context, in *proto.GetIndicesRequest, opts ...grpc.CallOption) (*proto.GetIndicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIndices", varargs...)
	ret0, _ := ret[0].(*proto.GetIndicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndices indicates an expected call of GetIndices.
func (mr *MockGatewayClientMockRecorder) GetIndices(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndices", reflect.TypeOf((*MockGatewayClient)(nil).GetIndices), varargs...)
}

// GetMovieByID mocks base method.
func (m *MockGatewayClient) GetMovieByID(ctx context.Context, in *proto.GetMovieByIDRequest, opts ...grpc.CallOption) (*proto.GetMovieByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovieByID", varargs...)
	ret0, _ := ret[0].(*proto.GetMovieByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByID indicates an expected call of GetMovieByID.
func (mr *MockGatewayClientMockRecorder) GetMovieByID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByID", reflect.TypeOf((*MockGatewayClient)(nil).GetMovieByID), varargs...)
}

// SearchMovie mocks base method.
func (m *MockGatewayClient) SearchMovie(ctx context.Context, in *proto.SearchMovieRequest, opts ...grpc.CallOption) (*proto.SearchMovieResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMovie", varargs...)
	ret0, _ := ret[0].(*proto.SearchMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockGatewayClientMockRecorder) SearchMovie(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockGatewayClient)(nil).SearchMovie), varargs...)
}

// UpdateMovieData mocks base method.
func (m *MockGatewayClient) UpdateMovieData(ctx context.Context, in *proto.UpdateMovieDataRequest, opts ...grpc.CallOption) (*proto.UpdateMovieDataResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMovieData", varargs...)
	ret0, _ := ret[0].(*proto.UpdateMovieDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMovieData indicates an expected call of UpdateMovieData.
func (mr *MockGatewayClientMockRecorder) UpdateMovieData(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovieData", reflect.TypeOf((*MockGatewayClient)(nil).UpdateMovieData), varargs...)
}

// MockGatewayServer is a mock of GatewayServer interface.
type MockGatewayServer struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayServerMockRecorder
}

// MockGatewayServerMockRecorder is the mock recorder for MockGatewayServer.
type MockGatewayServerMockRecorder struct {
	mock *MockGatewayServer
}

// NewMockGatewayServer creates a new mock instance.
func NewMockGatewayServer(ctrl *gomock.Controller) *MockGatewayServer {
	mock := &MockGatewayServer{ctrl: ctrl}
	mock.recorder = &MockGatewayServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGatewayServer) EXPECT() *MockGatewayServerMockRecorder {
	return m.recorder
}

// BulkAPI mocks base method.
func (m *MockGatewayServer) BulkAPI(arg0 context.Context, arg1 *proto.BulkAPIRequest) (*proto.BulkAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkAPI", arg0, arg1)
	ret0, _ := ret[0].(*proto.BulkAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkAPI indicates an expected call of BulkAPI.
func (mr *MockGatewayServerMockRecorder) BulkAPI(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkAPI", reflect.TypeOf((*MockGatewayServer)(nil).BulkAPI), arg0, arg1)
}

// CreateNewMovie mocks base method.
func (m *MockGatewayServer) CreateNewMovie(arg0 context.Context, arg1 *proto.CreateNewMovieRequest) (*proto.CreateNewMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewMovie", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreateNewMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewMovie indicates an expected call of CreateNewMovie.
func (mr *MockGatewayServerMockRecorder) CreateNewMovie(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewMovie", reflect.TypeOf((*MockGatewayServer)(nil).CreateNewMovie), arg0, arg1)
}

// DeleteMovie mocks base method.
func (m *MockGatewayServer) DeleteMovie(arg0 context.Context, arg1 *proto.DeleteMovieRequest) (*proto.DeleteMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", arg0, arg1)
	ret0, _ := ret[0].(*proto.DeleteMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockGatewayServerMockRecorder) DeleteMovie(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockGatewayServer)(nil).DeleteMovie), arg0, arg1)
}

// GetAllMovie mocks base method.
func (m *MockGatewayServer) GetAllMovie(arg0 context.Context, arg1 *proto.GetAllMovieRequest) (*proto.GetAllMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMovie", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetAllMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMovie indicates an expected call of GetAllMovie.
func (mr *MockGatewayServerMockRecorder) GetAllMovie(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMovie", reflect.TypeOf((*MockGatewayServer)(nil).GetAllMovie), arg0, arg1)
}

// GetIndices mocks base method.
func (m *MockGatewayServer) GetIndices(arg0 context.Context, arg1 *proto.GetIndicesRequest) (*proto.GetIndicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndices", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetIndicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndices indicates an expected call of GetIndices.
func (mr *MockGatewayServerMockRecorder) GetIndices(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndices", reflect.TypeOf((*MockGatewayServer)(nil).GetIndices), arg0, arg1)
}

// GetMovieByID mocks base method.
func (m *MockGatewayServer) GetMovieByID(arg0 context.Context, arg1 *proto.GetMovieByIDRequest) (*proto.GetMovieByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByID", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetMovieByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByID indicates an expected call of GetMovieByID.
func (mr *MockGatewayServerMockRecorder) GetMovieByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByID", reflect.TypeOf((*MockGatewayServer)(nil).GetMovieByID), arg0, arg1)
}

// SearchMovie mocks base method.
func (m *MockGatewayServer) SearchMovie(arg0 context.Context, arg1 *proto.SearchMovieRequest) (*proto.SearchMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovie", arg0, arg1)
	ret0, _ := ret[0].(*proto.SearchMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockGatewayServerMockRecorder) SearchMovie(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockGatewayServer)(nil).SearchMovie), arg0, arg1)
}

// UpdateMovieData mocks base method.
func (m *MockGatewayServer) UpdateMovieData(arg0 context.Context, arg1 *proto.UpdateMovieDataRequest) (*proto.UpdateMovieDataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovieData", arg0, arg1)
	ret0, _ := ret[0].(*proto.UpdateMovieDataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMovieData indicates an expected call of UpdateMovieData.
func (mr *MockGatewayServerMockRecorder) UpdateMovieData(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovieData", reflect.TypeOf((*MockGatewayServer)(nil).UpdateMovieData), arg0, arg1)
}

// mustEmbedUnimplementedGatewayServer mocks base method.
func (m *MockGatewayServer) mustEmbedUnimplementedGatewayServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGatewayServer")
}

// mustEmbedUnimplementedGatewayServer indicates an expected call of mustEmbedUnimplementedGatewayServer.
func (mr *MockGatewayServerMockRecorder) mustEmbedUnimplementedGatewayServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGatewayServer", reflect.TypeOf((*MockGatewayServer)(nil).mustEmbedUnimplementedGatewayServer))
}

// MockUnsafeGatewayServer is a mock of UnsafeGatewayServer interface.
type MockUnsafeGatewayServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeGatewayServerMockRecorder
}

// MockUnsafeGatewayServerMockRecorder is the mock recorder for MockUnsafeGatewayServer.
type MockUnsafeGatewayServerMockRecorder struct {
	mock *MockUnsafeGatewayServer
}

// NewMockUnsafeGatewayServer creates a new mock instance.
func NewMockUnsafeGatewayServer(ctrl *gomock.Controller) *MockUnsafeGatewayServer {
	mock := &MockUnsafeGatewayServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeGatewayServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeGatewayServer) EXPECT() *MockUnsafeGatewayServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedGatewayServer mocks base method.
func (m *MockUnsafeGatewayServer) mustEmbedUnimplementedGatewayServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedGatewayServer")
}

// mustEmbedUnimplementedGatewayServer indicates an expected call of mustEmbedUnimplementedGatewayServer.
func (mr *MockUnsafeGatewayServerMockRecorder) mustEmbedUnimplementedGatewayServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedGatewayServer", reflect.TypeOf((*MockUnsafeGatewayServer)(nil).mustEmbedUnimplementedGatewayServer))
}

// MockAggregatorClient is a mock of AggregatorClient interface.
type MockAggregatorClient struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorClientMockRecorder
}

// MockAggregatorClientMockRecorder is the mock recorder for MockAggregatorClient.
type MockAggregatorClientMockRecorder struct {
	mock *MockAggregatorClient
}

// NewMockAggregatorClient creates a new mock instance.
func NewMockAggregatorClient(ctrl *gomock.Controller) *MockAggregatorClient {
	mock := &MockAggregatorClient{ctrl: ctrl}
	mock.recorder = &MockAggregatorClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregatorClient) EXPECT() *MockAggregatorClientMockRecorder {
	return m.recorder
}

// CreateMoviePostgres mocks base method.
func (m *MockAggregatorClient) CreateMoviePostgres(ctx context.Context, in *proto.CreateMoviePostgresRequest, opts ...grpc.CallOption) (*proto.CreateMoviePostgresResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMoviePostgres", varargs...)
	ret0, _ := ret[0].(*proto.CreateMoviePostgresResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMoviePostgres indicates an expected call of CreateMoviePostgres.
func (mr *MockAggregatorClientMockRecorder) CreateMoviePostgres(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMoviePostgres", reflect.TypeOf((*MockAggregatorClient)(nil).CreateMoviePostgres), varargs...)
}

// GetCache mocks base method.
func (m *MockAggregatorClient) GetCache(ctx context.Context, in *proto.GetCacheRequest, opts ...grpc.CallOption) (*proto.GetCacheResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCache", varargs...)
	ret0, _ := ret[0].(*proto.GetCacheResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockAggregatorClientMockRecorder) GetCache(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockAggregatorClient)(nil).GetCache), varargs...)
}

// SearchMovieAggregator mocks base method.
func (m *MockAggregatorClient) SearchMovieAggregator(ctx context.Context, in *proto.SearchMovieAggregatorRequest, opts ...grpc.CallOption) (*proto.SearchMovieAggregatorResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchMovieAggregator", varargs...)
	ret0, _ := ret[0].(*proto.SearchMovieAggregatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovieAggregator indicates an expected call of SearchMovieAggregator.
func (mr *MockAggregatorClientMockRecorder) SearchMovieAggregator(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovieAggregator", reflect.TypeOf((*MockAggregatorClient)(nil).SearchMovieAggregator), varargs...)
}

// SetCache mocks base method.
func (m *MockAggregatorClient) SetCache(ctx context.Context, in *proto.SetCacheRequest, opts ...grpc.CallOption) (*proto.SetCacheResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetCache", varargs...)
	ret0, _ := ret[0].(*proto.SetCacheResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCache indicates an expected call of SetCache.
func (mr *MockAggregatorClientMockRecorder) SetCache(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCache", reflect.TypeOf((*MockAggregatorClient)(nil).SetCache), varargs...)
}

// MockAggregatorServer is a mock of AggregatorServer interface.
type MockAggregatorServer struct {
	ctrl     *gomock.Controller
	recorder *MockAggregatorServerMockRecorder
}

// MockAggregatorServerMockRecorder is the mock recorder for MockAggregatorServer.
type MockAggregatorServerMockRecorder struct {
	mock *MockAggregatorServer
}

// NewMockAggregatorServer creates a new mock instance.
func NewMockAggregatorServer(ctrl *gomock.Controller) *MockAggregatorServer {
	mock := &MockAggregatorServer{ctrl: ctrl}
	mock.recorder = &MockAggregatorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAggregatorServer) EXPECT() *MockAggregatorServerMockRecorder {
	return m.recorder
}

// CreateMoviePostgres mocks base method.
func (m *MockAggregatorServer) CreateMoviePostgres(arg0 context.Context, arg1 *proto.CreateMoviePostgresRequest) (*proto.CreateMoviePostgresResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMoviePostgres", arg0, arg1)
	ret0, _ := ret[0].(*proto.CreateMoviePostgresResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMoviePostgres indicates an expected call of CreateMoviePostgres.
func (mr *MockAggregatorServerMockRecorder) CreateMoviePostgres(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMoviePostgres", reflect.TypeOf((*MockAggregatorServer)(nil).CreateMoviePostgres), arg0, arg1)
}

// GetCache mocks base method.
func (m *MockAggregatorServer) GetCache(arg0 context.Context, arg1 *proto.GetCacheRequest) (*proto.GetCacheResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetCacheResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockAggregatorServerMockRecorder) GetCache(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockAggregatorServer)(nil).GetCache), arg0, arg1)
}

// SearchMovieAggregator mocks base method.
func (m *MockAggregatorServer) SearchMovieAggregator(arg0 context.Context, arg1 *proto.SearchMovieAggregatorRequest) (*proto.SearchMovieAggregatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovieAggregator", arg0, arg1)
	ret0, _ := ret[0].(*proto.SearchMovieAggregatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovieAggregator indicates an expected call of SearchMovieAggregator.
func (mr *MockAggregatorServerMockRecorder) SearchMovieAggregator(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovieAggregator", reflect.TypeOf((*MockAggregatorServer)(nil).SearchMovieAggregator), arg0, arg1)
}

// SetCache mocks base method.
func (m *MockAggregatorServer) SetCache(arg0 context.Context, arg1 *proto.SetCacheRequest) (*proto.SetCacheResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCache", arg0, arg1)
	ret0, _ := ret[0].(*proto.SetCacheResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCache indicates an expected call of SetCache.
func (mr *MockAggregatorServerMockRecorder) SetCache(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCache", reflect.TypeOf((*MockAggregatorServer)(nil).SetCache), arg0, arg1)
}

// mustEmbedUnimplementedAggregatorServer mocks base method.
func (m *MockAggregatorServer) mustEmbedUnimplementedAggregatorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAggregatorServer")
}

// mustEmbedUnimplementedAggregatorServer indicates an expected call of mustEmbedUnimplementedAggregatorServer.
func (mr *MockAggregatorServerMockRecorder) mustEmbedUnimplementedAggregatorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAggregatorServer", reflect.TypeOf((*MockAggregatorServer)(nil).mustEmbedUnimplementedAggregatorServer))
}

// MockUnsafeAggregatorServer is a mock of UnsafeAggregatorServer interface.
type MockUnsafeAggregatorServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeAggregatorServerMockRecorder
}

// MockUnsafeAggregatorServerMockRecorder is the mock recorder for MockUnsafeAggregatorServer.
type MockUnsafeAggregatorServerMockRecorder struct {
	mock *MockUnsafeAggregatorServer
}

// NewMockUnsafeAggregatorServer creates a new mock instance.
func NewMockUnsafeAggregatorServer(ctrl *gomock.Controller) *MockUnsafeAggregatorServer {
	mock := &MockUnsafeAggregatorServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeAggregatorServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAggregatorServer) EXPECT() *MockUnsafeAggregatorServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedAggregatorServer mocks base method.
func (m *MockUnsafeAggregatorServer) mustEmbedUnimplementedAggregatorServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedAggregatorServer")
}

// mustEmbedUnimplementedAggregatorServer indicates an expected call of mustEmbedUnimplementedAggregatorServer.
func (mr *MockUnsafeAggregatorServerMockRecorder) mustEmbedUnimplementedAggregatorServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAggregatorServer", reflect.TypeOf((*MockUnsafeAggregatorServer)(nil).mustEmbedUnimplementedAggregatorServer))
}
