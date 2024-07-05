// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/api/config/config.proto

package config

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

// NodeMetricClient is the client API for NodeMetric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeMetricClient interface {
	ReceiveNodeMetric(ctx context.Context, in *NodeMetricRequest, opts ...grpc.CallOption) (*MetricResponse, error)
}

type nodeMetricClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeMetricClient(cc grpc.ClientConnInterface) NodeMetricClient {
	return &nodeMetricClient{cc}
}

func (c *nodeMetricClient) ReceiveNodeMetric(ctx context.Context, in *NodeMetricRequest, opts ...grpc.CallOption) (*MetricResponse, error) {
	out := new(MetricResponse)
	err := c.cc.Invoke(ctx, "/config.NodeMetric/receiveNodeMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeMetricServer is the server API for NodeMetric service.
// All implementations must embed UnimplementedNodeMetricServer
// for forward compatibility
type NodeMetricServer interface {
	ReceiveNodeMetric(context.Context, *NodeMetricRequest) (*MetricResponse, error)
	mustEmbedUnimplementedNodeMetricServer()
}

// UnimplementedNodeMetricServer must be embedded to have forward compatible implementations.
type UnimplementedNodeMetricServer struct {
}

func (UnimplementedNodeMetricServer) ReceiveNodeMetric(context.Context, *NodeMetricRequest) (*MetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveNodeMetric not implemented")
}
func (UnimplementedNodeMetricServer) mustEmbedUnimplementedNodeMetricServer() {}

// UnsafeNodeMetricServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeMetricServer will
// result in compilation errors.
type UnsafeNodeMetricServer interface {
	mustEmbedUnimplementedNodeMetricServer()
}

func RegisterNodeMetricServer(s grpc.ServiceRegistrar, srv NodeMetricServer) {
	s.RegisterService(&NodeMetric_ServiceDesc, srv)
}

func _NodeMetric_ReceiveNodeMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeMetricServer).ReceiveNodeMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.NodeMetric/receiveNodeMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeMetricServer).ReceiveNodeMetric(ctx, req.(*NodeMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeMetric_ServiceDesc is the grpc.ServiceDesc for NodeMetric service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeMetric_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.NodeMetric",
	HandlerType: (*NodeMetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "receiveNodeMetric",
			Handler:    _NodeMetric_ReceiveNodeMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/config/config.proto",
}

// CSDMetricClient is the client API for CSDMetric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CSDMetricClient interface {
	ReceiveCSDMetric(ctx context.Context, in *CSDMetricRequest, opts ...grpc.CallOption) (*MetricResponse, error)
}

type cSDMetricClient struct {
	cc grpc.ClientConnInterface
}

func NewCSDMetricClient(cc grpc.ClientConnInterface) CSDMetricClient {
	return &cSDMetricClient{cc}
}

func (c *cSDMetricClient) ReceiveCSDMetric(ctx context.Context, in *CSDMetricRequest, opts ...grpc.CallOption) (*MetricResponse, error) {
	out := new(MetricResponse)
	err := c.cc.Invoke(ctx, "/config.CSDMetric/receiveCSDMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSDMetricServer is the server API for CSDMetric service.
// All implementations must embed UnimplementedCSDMetricServer
// for forward compatibility
type CSDMetricServer interface {
	ReceiveCSDMetric(context.Context, *CSDMetricRequest) (*MetricResponse, error)
	mustEmbedUnimplementedCSDMetricServer()
}

// UnimplementedCSDMetricServer must be embedded to have forward compatible implementations.
type UnimplementedCSDMetricServer struct {
}

func (UnimplementedCSDMetricServer) ReceiveCSDMetric(context.Context, *CSDMetricRequest) (*MetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveCSDMetric not implemented")
}
func (UnimplementedCSDMetricServer) mustEmbedUnimplementedCSDMetricServer() {}

// UnsafeCSDMetricServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CSDMetricServer will
// result in compilation errors.
type UnsafeCSDMetricServer interface {
	mustEmbedUnimplementedCSDMetricServer()
}

func RegisterCSDMetricServer(s grpc.ServiceRegistrar, srv CSDMetricServer) {
	s.RegisterService(&CSDMetric_ServiceDesc, srv)
}

func _CSDMetric_ReceiveCSDMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSDMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSDMetricServer).ReceiveCSDMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.CSDMetric/receiveCSDMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSDMetricServer).ReceiveCSDMetric(ctx, req.(*CSDMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CSDMetric_ServiceDesc is the grpc.ServiceDesc for CSDMetric service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CSDMetric_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.CSDMetric",
	HandlerType: (*CSDMetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "receiveCSDMetric",
			Handler:    _CSDMetric_ReceiveCSDMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/config/config.proto",
}
