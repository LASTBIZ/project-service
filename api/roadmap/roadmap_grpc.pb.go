// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.5
// source: api/roadmap/roadmap.proto

package roadmap

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Roadmap_AddRoadMap_FullMethodName    = "/api.roadmap.Roadmap/AddRoadMap"
	Roadmap_UpdateRoadMap_FullMethodName = "/api.roadmap.Roadmap/UpdateRoadMap"
	Roadmap_RemoveRoadMap_FullMethodName = "/api.roadmap.Roadmap/RemoveRoadMap"
	Roadmap_ListRoadmap_FullMethodName   = "/api.roadmap.Roadmap/ListRoadmap"
)

// RoadmapClient is the client API for Roadmap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoadmapClient interface {
	AddRoadMap(ctx context.Context, in *CreateRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateRoadMap(ctx context.Context, in *UpdateRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveRoadMap(ctx context.Context, in *DeleteRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListRoadmap(ctx context.Context, in *ListRoadmapRequest, opts ...grpc.CallOption) (*ListRoadmapReply, error)
}

type roadmapClient struct {
	cc grpc.ClientConnInterface
}

func NewRoadmapClient(cc grpc.ClientConnInterface) RoadmapClient {
	return &roadmapClient{cc}
}

func (c *roadmapClient) AddRoadMap(ctx context.Context, in *CreateRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Roadmap_AddRoadMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadmapClient) UpdateRoadMap(ctx context.Context, in *UpdateRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Roadmap_UpdateRoadMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadmapClient) RemoveRoadMap(ctx context.Context, in *DeleteRoadmapRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Roadmap_RemoveRoadMap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roadmapClient) ListRoadmap(ctx context.Context, in *ListRoadmapRequest, opts ...grpc.CallOption) (*ListRoadmapReply, error) {
	out := new(ListRoadmapReply)
	err := c.cc.Invoke(ctx, Roadmap_ListRoadmap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoadmapServer is the server API for Roadmap service.
// All implementations must embed UnimplementedRoadmapServer
// for forward compatibility
type RoadmapServer interface {
	AddRoadMap(context.Context, *CreateRoadmapRequest) (*emptypb.Empty, error)
	UpdateRoadMap(context.Context, *UpdateRoadmapRequest) (*emptypb.Empty, error)
	RemoveRoadMap(context.Context, *DeleteRoadmapRequest) (*emptypb.Empty, error)
	ListRoadmap(context.Context, *ListRoadmapRequest) (*ListRoadmapReply, error)
	mustEmbedUnimplementedRoadmapServer()
}

// UnimplementedRoadmapServer must be embedded to have forward compatible implementations.
type UnimplementedRoadmapServer struct {
}

func (UnimplementedRoadmapServer) AddRoadMap(context.Context, *CreateRoadmapRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoadMap not implemented")
}
func (UnimplementedRoadmapServer) UpdateRoadMap(context.Context, *UpdateRoadmapRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoadMap not implemented")
}
func (UnimplementedRoadmapServer) RemoveRoadMap(context.Context, *DeleteRoadmapRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoadMap not implemented")
}
func (UnimplementedRoadmapServer) ListRoadmap(context.Context, *ListRoadmapRequest) (*ListRoadmapReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoadmap not implemented")
}
func (UnimplementedRoadmapServer) mustEmbedUnimplementedRoadmapServer() {}

// UnsafeRoadmapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoadmapServer will
// result in compilation errors.
type UnsafeRoadmapServer interface {
	mustEmbedUnimplementedRoadmapServer()
}

func RegisterRoadmapServer(s grpc.ServiceRegistrar, srv RoadmapServer) {
	s.RegisterService(&Roadmap_ServiceDesc, srv)
}

func _Roadmap_AddRoadMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadmapServer).AddRoadMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roadmap_AddRoadMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadmapServer).AddRoadMap(ctx, req.(*CreateRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roadmap_UpdateRoadMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadmapServer).UpdateRoadMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roadmap_UpdateRoadMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadmapServer).UpdateRoadMap(ctx, req.(*UpdateRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roadmap_RemoveRoadMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadmapServer).RemoveRoadMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roadmap_RemoveRoadMap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadmapServer).RemoveRoadMap(ctx, req.(*DeleteRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roadmap_ListRoadmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoadmapServer).ListRoadmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roadmap_ListRoadmap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoadmapServer).ListRoadmap(ctx, req.(*ListRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Roadmap_ServiceDesc is the grpc.ServiceDesc for Roadmap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roadmap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.roadmap.Roadmap",
	HandlerType: (*RoadmapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRoadMap",
			Handler:    _Roadmap_AddRoadMap_Handler,
		},
		{
			MethodName: "UpdateRoadMap",
			Handler:    _Roadmap_UpdateRoadMap_Handler,
		},
		{
			MethodName: "RemoveRoadMap",
			Handler:    _Roadmap_RemoveRoadMap_Handler,
		},
		{
			MethodName: "ListRoadmap",
			Handler:    _Roadmap_ListRoadmap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/roadmap/roadmap.proto",
}
