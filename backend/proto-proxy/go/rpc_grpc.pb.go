// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rpc.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Gateway_GetAllMovie_FullMethodName     = "/proxy_proto.Gateway/GetAllMovie"
	Gateway_CreateNewMovie_FullMethodName  = "/proxy_proto.Gateway/CreateNewMovie"
	Gateway_GetIndices_FullMethodName      = "/proxy_proto.Gateway/GetIndices"
	Gateway_GetMovieByID_FullMethodName    = "/proxy_proto.Gateway/GetMovieByID"
	Gateway_SearchMovie_FullMethodName     = "/proxy_proto.Gateway/SearchMovie"
	Gateway_UpdateMovieData_FullMethodName = "/proxy_proto.Gateway/UpdateMovieData"
	Gateway_DeleteMovie_FullMethodName     = "/proxy_proto.Gateway/DeleteMovie"
	Gateway_BulkAPI_FullMethodName         = "/proxy_proto.Gateway/BulkAPI"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetAllMovie(ctx context.Context, in *GetAllMovieRequest, opts ...grpc.CallOption) (*GetAllMovieResponse, error)
	CreateNewMovie(ctx context.Context, in *CreateNewMovieRequest, opts ...grpc.CallOption) (*CreateNewMovieResponse, error)
	GetIndices(ctx context.Context, in *GetIndicesRequest, opts ...grpc.CallOption) (*GetIndicesResponse, error)
	GetMovieByID(ctx context.Context, in *GetMovieByIDRequest, opts ...grpc.CallOption) (*GetMovieByIDResponse, error)
	SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*SearchMovieResponse, error)
	UpdateMovieData(ctx context.Context, in *UpdateMovieDataRequest, opts ...grpc.CallOption) (*UpdateMovieDataResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error)
	BulkAPI(ctx context.Context, in *BulkAPIRequest, opts ...grpc.CallOption) (*BulkAPIResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetAllMovie(ctx context.Context, in *GetAllMovieRequest, opts ...grpc.CallOption) (*GetAllMovieResponse, error) {
	out := new(GetAllMovieResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAllMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateNewMovie(ctx context.Context, in *CreateNewMovieRequest, opts ...grpc.CallOption) (*CreateNewMovieResponse, error) {
	out := new(CreateNewMovieResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateNewMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetIndices(ctx context.Context, in *GetIndicesRequest, opts ...grpc.CallOption) (*GetIndicesResponse, error) {
	out := new(GetIndicesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetIndices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetMovieByID(ctx context.Context, in *GetMovieByIDRequest, opts ...grpc.CallOption) (*GetMovieByIDResponse, error) {
	out := new(GetMovieByIDResponse)
	err := c.cc.Invoke(ctx, Gateway_GetMovieByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*SearchMovieResponse, error) {
	out := new(SearchMovieResponse)
	err := c.cc.Invoke(ctx, Gateway_SearchMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateMovieData(ctx context.Context, in *UpdateMovieDataRequest, opts ...grpc.CallOption) (*UpdateMovieDataResponse, error) {
	out := new(UpdateMovieDataResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateMovieData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*DeleteMovieResponse, error) {
	out := new(DeleteMovieResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) BulkAPI(ctx context.Context, in *BulkAPIRequest, opts ...grpc.CallOption) (*BulkAPIResponse, error) {
	out := new(BulkAPIResponse)
	err := c.cc.Invoke(ctx, Gateway_BulkAPI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetAllMovie(context.Context, *GetAllMovieRequest) (*GetAllMovieResponse, error)
	CreateNewMovie(context.Context, *CreateNewMovieRequest) (*CreateNewMovieResponse, error)
	GetIndices(context.Context, *GetIndicesRequest) (*GetIndicesResponse, error)
	GetMovieByID(context.Context, *GetMovieByIDRequest) (*GetMovieByIDResponse, error)
	SearchMovie(context.Context, *SearchMovieRequest) (*SearchMovieResponse, error)
	UpdateMovieData(context.Context, *UpdateMovieDataRequest) (*UpdateMovieDataResponse, error)
	DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error)
	BulkAPI(context.Context, *BulkAPIRequest) (*BulkAPIResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (UnimplementedGatewayServer) GetAllMovie(context.Context, *GetAllMovieRequest) (*GetAllMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMovie not implemented")
}
func (UnimplementedGatewayServer) CreateNewMovie(context.Context, *CreateNewMovieRequest) (*CreateNewMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewMovie not implemented")
}
func (UnimplementedGatewayServer) GetIndices(context.Context, *GetIndicesRequest) (*GetIndicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndices not implemented")
}
func (UnimplementedGatewayServer) GetMovieByID(context.Context, *GetMovieByIDRequest) (*GetMovieByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByID not implemented")
}
func (UnimplementedGatewayServer) SearchMovie(context.Context, *SearchMovieRequest) (*SearchMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovie not implemented")
}
func (UnimplementedGatewayServer) UpdateMovieData(context.Context, *UpdateMovieDataRequest) (*UpdateMovieDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovieData not implemented")
}
func (UnimplementedGatewayServer) DeleteMovie(context.Context, *DeleteMovieRequest) (*DeleteMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedGatewayServer) BulkAPI(context.Context, *BulkAPIRequest) (*BulkAPIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkAPI not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_GetAllMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAllMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAllMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAllMovie(ctx, req.(*GetAllMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateNewMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateNewMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateNewMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateNewMovie(ctx, req.(*CreateNewMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetIndices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetIndices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetIndices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetIndices(ctx, req.(*GetIndicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetMovieByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetMovieByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetMovieByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetMovieByID(ctx, req.(*GetMovieByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_SearchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).SearchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_SearchMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).SearchMovie(ctx, req.(*SearchMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateMovieData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateMovieData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateMovieData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateMovieData(ctx, req.(*UpdateMovieDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_BulkAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkAPIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).BulkAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_BulkAPI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).BulkAPI(ctx, req.(*BulkAPIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy_proto.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMovie",
			Handler:    _Gateway_GetAllMovie_Handler,
		},
		{
			MethodName: "CreateNewMovie",
			Handler:    _Gateway_CreateNewMovie_Handler,
		},
		{
			MethodName: "GetIndices",
			Handler:    _Gateway_GetIndices_Handler,
		},
		{
			MethodName: "GetMovieByID",
			Handler:    _Gateway_GetMovieByID_Handler,
		},
		{
			MethodName: "SearchMovie",
			Handler:    _Gateway_SearchMovie_Handler,
		},
		{
			MethodName: "UpdateMovieData",
			Handler:    _Gateway_UpdateMovieData_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _Gateway_DeleteMovie_Handler,
		},
		{
			MethodName: "BulkAPI",
			Handler:    _Gateway_BulkAPI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

const (
	Aggregator_SearchMovieAggregator_FullMethodName = "/proxy_proto.Aggregator/SearchMovieAggregator"
	Aggregator_SetCache_FullMethodName              = "/proxy_proto.Aggregator/SetCache"
	Aggregator_GetCache_FullMethodName              = "/proxy_proto.Aggregator/GetCache"
	Aggregator_CreateMoviePostgres_FullMethodName   = "/proxy_proto.Aggregator/CreateMoviePostgres"
	Aggregator_SearchMoviePostgres_FullMethodName   = "/proxy_proto.Aggregator/SearchMoviePostgres"
)

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregatorClient interface {
	SearchMovieAggregator(ctx context.Context, in *SearchMovieAggregatorRequest, opts ...grpc.CallOption) (*SearchMovieAggregatorResponse, error)
	SetCache(ctx context.Context, in *SetCacheRequest, opts ...grpc.CallOption) (*SetCacheResponse, error)
	GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheResponse, error)
	CreateMoviePostgres(ctx context.Context, in *CreateMoviePostgresRequest, opts ...grpc.CallOption) (*CreateMoviePostgresResponse, error)
	SearchMoviePostgres(ctx context.Context, in *SearchMoviePostgresRequest, opts ...grpc.CallOption) (*SearchMoviePostgresResponse, error)
}

type aggregatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregatorClient(cc grpc.ClientConnInterface) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) SearchMovieAggregator(ctx context.Context, in *SearchMovieAggregatorRequest, opts ...grpc.CallOption) (*SearchMovieAggregatorResponse, error) {
	out := new(SearchMovieAggregatorResponse)
	err := c.cc.Invoke(ctx, Aggregator_SearchMovieAggregator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) SetCache(ctx context.Context, in *SetCacheRequest, opts ...grpc.CallOption) (*SetCacheResponse, error) {
	out := new(SetCacheResponse)
	err := c.cc.Invoke(ctx, Aggregator_SetCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) GetCache(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheResponse, error) {
	out := new(GetCacheResponse)
	err := c.cc.Invoke(ctx, Aggregator_GetCache_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) CreateMoviePostgres(ctx context.Context, in *CreateMoviePostgresRequest, opts ...grpc.CallOption) (*CreateMoviePostgresResponse, error) {
	out := new(CreateMoviePostgresResponse)
	err := c.cc.Invoke(ctx, Aggregator_CreateMoviePostgres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) SearchMoviePostgres(ctx context.Context, in *SearchMoviePostgresRequest, opts ...grpc.CallOption) (*SearchMoviePostgresResponse, error) {
	out := new(SearchMoviePostgresResponse)
	err := c.cc.Invoke(ctx, Aggregator_SearchMoviePostgres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
// All implementations must embed UnimplementedAggregatorServer
// for forward compatibility
type AggregatorServer interface {
	SearchMovieAggregator(context.Context, *SearchMovieAggregatorRequest) (*SearchMovieAggregatorResponse, error)
	SetCache(context.Context, *SetCacheRequest) (*SetCacheResponse, error)
	GetCache(context.Context, *GetCacheRequest) (*GetCacheResponse, error)
	CreateMoviePostgres(context.Context, *CreateMoviePostgresRequest) (*CreateMoviePostgresResponse, error)
	SearchMoviePostgres(context.Context, *SearchMoviePostgresRequest) (*SearchMoviePostgresResponse, error)
	mustEmbedUnimplementedAggregatorServer()
}

// UnimplementedAggregatorServer must be embedded to have forward compatible implementations.
type UnimplementedAggregatorServer struct {
}

func (UnimplementedAggregatorServer) SearchMovieAggregator(context.Context, *SearchMovieAggregatorRequest) (*SearchMovieAggregatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovieAggregator not implemented")
}
func (UnimplementedAggregatorServer) SetCache(context.Context, *SetCacheRequest) (*SetCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCache not implemented")
}
func (UnimplementedAggregatorServer) GetCache(context.Context, *GetCacheRequest) (*GetCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCache not implemented")
}
func (UnimplementedAggregatorServer) CreateMoviePostgres(context.Context, *CreateMoviePostgresRequest) (*CreateMoviePostgresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMoviePostgres not implemented")
}
func (UnimplementedAggregatorServer) SearchMoviePostgres(context.Context, *SearchMoviePostgresRequest) (*SearchMoviePostgresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMoviePostgres not implemented")
}
func (UnimplementedAggregatorServer) mustEmbedUnimplementedAggregatorServer() {}

// UnsafeAggregatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregatorServer will
// result in compilation errors.
type UnsafeAggregatorServer interface {
	mustEmbedUnimplementedAggregatorServer()
}

func RegisterAggregatorServer(s grpc.ServiceRegistrar, srv AggregatorServer) {
	s.RegisterService(&Aggregator_ServiceDesc, srv)
}

func _Aggregator_SearchMovieAggregator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMovieAggregatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).SearchMovieAggregator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_SearchMovieAggregator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).SearchMovieAggregator(ctx, req.(*SearchMovieAggregatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_SetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).SetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_SetCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).SetCache(ctx, req.(*SetCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_GetCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).GetCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_GetCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).GetCache(ctx, req.(*GetCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_CreateMoviePostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMoviePostgresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).CreateMoviePostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_CreateMoviePostgres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).CreateMoviePostgres(ctx, req.(*CreateMoviePostgresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_SearchMoviePostgres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMoviePostgresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).SearchMoviePostgres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aggregator_SearchMoviePostgres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).SearchMoviePostgres(ctx, req.(*SearchMoviePostgresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aggregator_ServiceDesc is the grpc.ServiceDesc for Aggregator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aggregator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy_proto.Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMovieAggregator",
			Handler:    _Aggregator_SearchMovieAggregator_Handler,
		},
		{
			MethodName: "SetCache",
			Handler:    _Aggregator_SetCache_Handler,
		},
		{
			MethodName: "GetCache",
			Handler:    _Aggregator_GetCache_Handler,
		},
		{
			MethodName: "CreateMoviePostgres",
			Handler:    _Aggregator_CreateMoviePostgres_Handler,
		},
		{
			MethodName: "SearchMoviePostgres",
			Handler:    _Aggregator_SearchMoviePostgres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
