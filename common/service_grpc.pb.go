// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: common/service.proto

package common

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NodeService_Start_FullMethodName                    = "/service.NodeService/Start"
	NodeService_Stop_FullMethodName                     = "/service.NodeService/Stop"
	NodeService_GetBaseInfo_FullMethodName              = "/service.NodeService/GetBaseInfo"
	NodeService_GetLogs_FullMethodName                  = "/service.NodeService/GetLogs"
	NodeService_GetSystemStats_FullMethodName           = "/service.NodeService/GetSystemStats"
	NodeService_GetBackendStats_FullMethodName          = "/service.NodeService/GetBackendStats"
	NodeService_GetOutboundsStats_FullMethodName        = "/service.NodeService/GetOutboundsStats"
	NodeService_GetOutboundStats_FullMethodName         = "/service.NodeService/GetOutboundStats"
	NodeService_GetInboundsStats_FullMethodName         = "/service.NodeService/GetInboundsStats"
	NodeService_GetInboundStats_FullMethodName          = "/service.NodeService/GetInboundStats"
	NodeService_GetUsersStats_FullMethodName            = "/service.NodeService/GetUsersStats"
	NodeService_GetUserStats_FullMethodName             = "/service.NodeService/GetUserStats"
	NodeService_GetUserOnlineStats_FullMethodName       = "/service.NodeService/GetUserOnlineStats"
	NodeService_GetUserOnlineIpListStats_FullMethodName = "/service.NodeService/GetUserOnlineIpListStats"
	NodeService_SyncUser_FullMethodName                 = "/service.NodeService/SyncUser"
	NodeService_SyncUsers_FullMethodName                = "/service.NodeService/SyncUsers"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for node management and connection
type NodeServiceClient interface {
	Start(ctx context.Context, in *Backend, opts ...grpc.CallOption) (*BaseInfoResponse, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetBaseInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseInfoResponse, error)
	GetLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Log], error)
	GetSystemStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemStatsResponse, error)
	GetBackendStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BackendStatsResponse, error)
	GetOutboundsStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetOutboundStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetInboundsStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetInboundStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetUsersStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetUserStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error)
	GetUserOnlineStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*OnlineStatResponse, error)
	GetUserOnlineIpListStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatsOnlineIpListResponse, error)
	SyncUser(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[User, Empty], error)
	SyncUsers(ctx context.Context, in *Users, opts ...grpc.CallOption) (*Empty, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Start(ctx context.Context, in *Backend, opts ...grpc.CallOption) (*BaseInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, NodeService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetBaseInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BaseInfoResponse)
	err := c.cc.Invoke(ctx, NodeService_GetBaseInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetLogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Log], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], NodeService_GetLogs_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, Log]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NodeService_GetLogsClient = grpc.ServerStreamingClient[Log]

func (c *nodeServiceClient) GetSystemStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SystemStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SystemStatsResponse)
	err := c.cc.Invoke(ctx, NodeService_GetSystemStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetBackendStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BackendStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BackendStatsResponse)
	err := c.cc.Invoke(ctx, NodeService_GetBackendStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetOutboundsStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetOutboundsStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetOutboundStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetOutboundStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetInboundsStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetInboundsStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetInboundStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetInboundStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetUsersStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetUsersStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetUserStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetUserStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetUserOnlineStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*OnlineStatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineStatResponse)
	err := c.cc.Invoke(ctx, NodeService_GetUserOnlineStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetUserOnlineIpListStats(ctx context.Context, in *StatRequest, opts ...grpc.CallOption) (*StatsOnlineIpListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatsOnlineIpListResponse)
	err := c.cc.Invoke(ctx, NodeService_GetUserOnlineIpListStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) SyncUser(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[User, Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[1], NodeService_SyncUser_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[User, Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NodeService_SyncUserClient = grpc.ClientStreamingClient[User, Empty]

func (c *nodeServiceClient) SyncUsers(ctx context.Context, in *Users, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, NodeService_SyncUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility.
//
// Service for node management and connection
type NodeServiceServer interface {
	Start(context.Context, *Backend) (*BaseInfoResponse, error)
	Stop(context.Context, *Empty) (*Empty, error)
	GetBaseInfo(context.Context, *Empty) (*BaseInfoResponse, error)
	GetLogs(*Empty, grpc.ServerStreamingServer[Log]) error
	GetSystemStats(context.Context, *Empty) (*SystemStatsResponse, error)
	GetBackendStats(context.Context, *Empty) (*BackendStatsResponse, error)
	GetOutboundsStats(context.Context, *StatRequest) (*StatResponse, error)
	GetOutboundStats(context.Context, *StatRequest) (*StatResponse, error)
	GetInboundsStats(context.Context, *StatRequest) (*StatResponse, error)
	GetInboundStats(context.Context, *StatRequest) (*StatResponse, error)
	GetUsersStats(context.Context, *StatRequest) (*StatResponse, error)
	GetUserStats(context.Context, *StatRequest) (*StatResponse, error)
	GetUserOnlineStats(context.Context, *StatRequest) (*OnlineStatResponse, error)
	GetUserOnlineIpListStats(context.Context, *StatRequest) (*StatsOnlineIpListResponse, error)
	SyncUser(grpc.ClientStreamingServer[User, Empty]) error
	SyncUsers(context.Context, *Users) (*Empty, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeServiceServer struct{}

func (UnimplementedNodeServiceServer) Start(context.Context, *Backend) (*BaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedNodeServiceServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedNodeServiceServer) GetBaseInfo(context.Context, *Empty) (*BaseInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaseInfo not implemented")
}
func (UnimplementedNodeServiceServer) GetLogs(*Empty, grpc.ServerStreamingServer[Log]) error {
	return status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedNodeServiceServer) GetSystemStats(context.Context, *Empty) (*SystemStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemStats not implemented")
}
func (UnimplementedNodeServiceServer) GetBackendStats(context.Context, *Empty) (*BackendStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackendStats not implemented")
}
func (UnimplementedNodeServiceServer) GetOutboundsStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutboundsStats not implemented")
}
func (UnimplementedNodeServiceServer) GetOutboundStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutboundStats not implemented")
}
func (UnimplementedNodeServiceServer) GetInboundsStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboundsStats not implemented")
}
func (UnimplementedNodeServiceServer) GetInboundStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInboundStats not implemented")
}
func (UnimplementedNodeServiceServer) GetUsersStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersStats not implemented")
}
func (UnimplementedNodeServiceServer) GetUserStats(context.Context, *StatRequest) (*StatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStats not implemented")
}
func (UnimplementedNodeServiceServer) GetUserOnlineStats(context.Context, *StatRequest) (*OnlineStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOnlineStats not implemented")
}
func (UnimplementedNodeServiceServer) GetUserOnlineIpListStats(context.Context, *StatRequest) (*StatsOnlineIpListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOnlineIpListStats not implemented")
}
func (UnimplementedNodeServiceServer) SyncUser(grpc.ClientStreamingServer[User, Empty]) error {
	return status.Errorf(codes.Unimplemented, "method SyncUser not implemented")
}
func (UnimplementedNodeServiceServer) SyncUsers(context.Context, *Users) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncUsers not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}
func (UnimplementedNodeServiceServer) testEmbeddedByValue()                     {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Backend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Start(ctx, req.(*Backend))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetBaseInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).GetLogs(m, &grpc.GenericServerStream[Empty, Log]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NodeService_GetLogsServer = grpc.ServerStreamingServer[Log]

func _NodeService_GetSystemStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetSystemStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetSystemStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetSystemStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetBackendStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetBackendStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetBackendStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetBackendStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetOutboundsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetOutboundsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetOutboundsStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetOutboundsStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetOutboundStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetOutboundStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetOutboundStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetOutboundStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetInboundsStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetInboundsStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetInboundsStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetInboundsStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetInboundStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetInboundStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetInboundStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetInboundStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetUsersStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetUsersStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetUsersStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetUsersStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetUserStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetUserStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetUserStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetUserStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetUserOnlineStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetUserOnlineStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetUserOnlineStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetUserOnlineStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetUserOnlineIpListStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetUserOnlineIpListStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetUserOnlineIpListStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetUserOnlineIpListStats(ctx, req.(*StatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_SyncUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServiceServer).SyncUser(&grpc.GenericServerStream[User, Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NodeService_SyncUserServer = grpc.ClientStreamingServer[User, Empty]

func _NodeService_SyncUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Users)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).SyncUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_SyncUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).SyncUsers(ctx, req.(*Users))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _NodeService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _NodeService_Stop_Handler,
		},
		{
			MethodName: "GetBaseInfo",
			Handler:    _NodeService_GetBaseInfo_Handler,
		},
		{
			MethodName: "GetSystemStats",
			Handler:    _NodeService_GetSystemStats_Handler,
		},
		{
			MethodName: "GetBackendStats",
			Handler:    _NodeService_GetBackendStats_Handler,
		},
		{
			MethodName: "GetOutboundsStats",
			Handler:    _NodeService_GetOutboundsStats_Handler,
		},
		{
			MethodName: "GetOutboundStats",
			Handler:    _NodeService_GetOutboundStats_Handler,
		},
		{
			MethodName: "GetInboundsStats",
			Handler:    _NodeService_GetInboundsStats_Handler,
		},
		{
			MethodName: "GetInboundStats",
			Handler:    _NodeService_GetInboundStats_Handler,
		},
		{
			MethodName: "GetUsersStats",
			Handler:    _NodeService_GetUsersStats_Handler,
		},
		{
			MethodName: "GetUserStats",
			Handler:    _NodeService_GetUserStats_Handler,
		},
		{
			MethodName: "GetUserOnlineStats",
			Handler:    _NodeService_GetUserOnlineStats_Handler,
		},
		{
			MethodName: "GetUserOnlineIpListStats",
			Handler:    _NodeService_GetUserOnlineIpListStats_Handler,
		},
		{
			MethodName: "SyncUsers",
			Handler:    _NodeService_SyncUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _NodeService_GetLogs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SyncUser",
			Handler:       _NodeService_SyncUser_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "common/service.proto",
}
