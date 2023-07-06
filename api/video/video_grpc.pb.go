// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/video/video.proto

package video

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
	Video_CreateVideo_FullMethodName       = "/api.video.Video/CreateVideo"
	Video_CreateScreenShoot_FullMethodName = "/api.video.Video/CreateScreenShoot"
	Video_UpdateScreenShoot_FullMethodName = "/api.video.Video/UpdateScreenShoot"
	Video_DeleteScreenShoot_FullMethodName = "/api.video.Video/DeleteScreenShoot"
	Video_UpdateVideo_FullMethodName       = "/api.video.Video/UpdateVideo"
	Video_DeleteVideo_FullMethodName       = "/api.video.Video/DeleteVideo"
)

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateScreenShoot(ctx context.Context, in *CreateScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateScreenShoot(ctx context.Context, in *UpdateScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteScreenShoot(ctx context.Context, in *DeleteScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_CreateVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) CreateScreenShoot(ctx context.Context, in *CreateScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_CreateScreenShoot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) UpdateScreenShoot(ctx context.Context, in *UpdateScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_UpdateScreenShoot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DeleteScreenShoot(ctx context.Context, in *DeleteScreenShotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_DeleteScreenShoot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_UpdateVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_DeleteVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	CreateVideo(context.Context, *CreateVideoRequest) (*emptypb.Empty, error)
	CreateScreenShoot(context.Context, *CreateScreenShotRequest) (*emptypb.Empty, error)
	UpdateScreenShoot(context.Context, *UpdateScreenShotRequest) (*emptypb.Empty, error)
	DeleteScreenShoot(context.Context, *DeleteScreenShotRequest) (*emptypb.Empty, error)
	UpdateVideo(context.Context, *UpdateVideoRequest) (*emptypb.Empty, error)
	DeleteVideo(context.Context, *DeleteVideoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) CreateVideo(context.Context, *CreateVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedVideoServer) CreateScreenShoot(context.Context, *CreateScreenShotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScreenShoot not implemented")
}
func (UnimplementedVideoServer) UpdateScreenShoot(context.Context, *UpdateScreenShotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScreenShoot not implemented")
}
func (UnimplementedVideoServer) DeleteScreenShoot(context.Context, *DeleteScreenShotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScreenShoot not implemented")
}
func (UnimplementedVideoServer) UpdateVideo(context.Context, *UpdateVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVideo not implemented")
}
func (UnimplementedVideoServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_CreateVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_CreateScreenShoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScreenShotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).CreateScreenShoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_CreateScreenShoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).CreateScreenShoot(ctx, req.(*CreateScreenShotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_UpdateScreenShoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScreenShotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UpdateScreenShoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_UpdateScreenShoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UpdateScreenShoot(ctx, req.(*UpdateScreenShotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DeleteScreenShoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScreenShotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DeleteScreenShoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_DeleteScreenShoot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DeleteScreenShoot(ctx, req.(*DeleteScreenShotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_UpdateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UpdateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_UpdateVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UpdateVideo(ctx, req.(*UpdateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_DeleteVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.video.Video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVideo",
			Handler:    _Video_CreateVideo_Handler,
		},
		{
			MethodName: "CreateScreenShoot",
			Handler:    _Video_CreateScreenShoot_Handler,
		},
		{
			MethodName: "UpdateScreenShoot",
			Handler:    _Video_UpdateScreenShoot_Handler,
		},
		{
			MethodName: "DeleteScreenShoot",
			Handler:    _Video_DeleteScreenShoot_Handler,
		},
		{
			MethodName: "UpdateVideo",
			Handler:    _Video_UpdateVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _Video_DeleteVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/video/video.proto",
}
