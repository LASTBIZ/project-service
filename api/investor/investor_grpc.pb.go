// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.5
// source: investor/investor.proto

package investor

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
	Investor_CreateInvestor_FullMethodName      = "/api.investor.Investor/CreateInvestor"
	Investor_DeleteInvestor_FullMethodName      = "/api.investor.Investor/DeleteInvestor"
	Investor_AddMoneyInvestor_FullMethodName    = "/api.investor.Investor/AddMoneyInvestor"
	Investor_RemoveMoneyInvestor_FullMethodName = "/api.investor.Investor/RemoveMoneyInvestor"
	Investor_SetMoneyInvestor_FullMethodName    = "/api.investor.Investor/SetMoneyInvestor"
	Investor_ListInvestor_FullMethodName        = "/api.investor.Investor/ListInvestor"
)

// InvestorClient is the client API for Investor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvestorClient interface {
	CreateInvestor(ctx context.Context, in *CreateInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteInvestor(ctx context.Context, in *DeleteInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddMoneyInvestor(ctx context.Context, in *AddMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveMoneyInvestor(ctx context.Context, in *RemoveMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetMoneyInvestor(ctx context.Context, in *SetMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListInvestor(ctx context.Context, in *ListInvestorRequest, opts ...grpc.CallOption) (*ListInvestorReply, error)
}

type investorClient struct {
	cc grpc.ClientConnInterface
}

func NewInvestorClient(cc grpc.ClientConnInterface) InvestorClient {
	return &investorClient{cc}
}

func (c *investorClient) CreateInvestor(ctx context.Context, in *CreateInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Investor_CreateInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorClient) DeleteInvestor(ctx context.Context, in *DeleteInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Investor_DeleteInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorClient) AddMoneyInvestor(ctx context.Context, in *AddMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Investor_AddMoneyInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorClient) RemoveMoneyInvestor(ctx context.Context, in *RemoveMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Investor_RemoveMoneyInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorClient) SetMoneyInvestor(ctx context.Context, in *SetMoneyInvestorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Investor_SetMoneyInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *investorClient) ListInvestor(ctx context.Context, in *ListInvestorRequest, opts ...grpc.CallOption) (*ListInvestorReply, error) {
	out := new(ListInvestorReply)
	err := c.cc.Invoke(ctx, Investor_ListInvestor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvestorServer is the server API for Investor service.
// All implementations must embed UnimplementedInvestorServer
// for forward compatibility
type InvestorServer interface {
	CreateInvestor(context.Context, *CreateInvestorRequest) (*emptypb.Empty, error)
	DeleteInvestor(context.Context, *DeleteInvestorRequest) (*emptypb.Empty, error)
	AddMoneyInvestor(context.Context, *AddMoneyInvestorRequest) (*emptypb.Empty, error)
	RemoveMoneyInvestor(context.Context, *RemoveMoneyInvestorRequest) (*emptypb.Empty, error)
	SetMoneyInvestor(context.Context, *SetMoneyInvestorRequest) (*emptypb.Empty, error)
	ListInvestor(context.Context, *ListInvestorRequest) (*ListInvestorReply, error)
	mustEmbedUnimplementedInvestorServer()
}

// UnimplementedInvestorServer must be embedded to have forward compatible implementations.
type UnimplementedInvestorServer struct {
}

func (UnimplementedInvestorServer) CreateInvestor(context.Context, *CreateInvestorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvestor not implemented")
}
func (UnimplementedInvestorServer) DeleteInvestor(context.Context, *DeleteInvestorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInvestor not implemented")
}
func (UnimplementedInvestorServer) AddMoneyInvestor(context.Context, *AddMoneyInvestorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMoneyInvestor not implemented")
}
func (UnimplementedInvestorServer) RemoveMoneyInvestor(context.Context, *RemoveMoneyInvestorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMoneyInvestor not implemented")
}
func (UnimplementedInvestorServer) SetMoneyInvestor(context.Context, *SetMoneyInvestorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMoneyInvestor not implemented")
}
func (UnimplementedInvestorServer) ListInvestor(context.Context, *ListInvestorRequest) (*ListInvestorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInvestor not implemented")
}
func (UnimplementedInvestorServer) mustEmbedUnimplementedInvestorServer() {}

// UnsafeInvestorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvestorServer will
// result in compilation errors.
type UnsafeInvestorServer interface {
	mustEmbedUnimplementedInvestorServer()
}

func RegisterInvestorServer(s grpc.ServiceRegistrar, srv InvestorServer) {
	s.RegisterService(&Investor_ServiceDesc, srv)
}

func _Investor_CreateInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).CreateInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_CreateInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).CreateInvestor(ctx, req.(*CreateInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Investor_DeleteInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).DeleteInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_DeleteInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).DeleteInvestor(ctx, req.(*DeleteInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Investor_AddMoneyInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMoneyInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).AddMoneyInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_AddMoneyInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).AddMoneyInvestor(ctx, req.(*AddMoneyInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Investor_RemoveMoneyInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMoneyInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).RemoveMoneyInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_RemoveMoneyInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).RemoveMoneyInvestor(ctx, req.(*RemoveMoneyInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Investor_SetMoneyInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMoneyInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).SetMoneyInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_SetMoneyInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).SetMoneyInvestor(ctx, req.(*SetMoneyInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Investor_ListInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvestorServer).ListInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Investor_ListInvestor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvestorServer).ListInvestor(ctx, req.(*ListInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Investor_ServiceDesc is the grpc.ServiceDesc for Investor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Investor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.investor.Investor",
	HandlerType: (*InvestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvestor",
			Handler:    _Investor_CreateInvestor_Handler,
		},
		{
			MethodName: "DeleteInvestor",
			Handler:    _Investor_DeleteInvestor_Handler,
		},
		{
			MethodName: "AddMoneyInvestor",
			Handler:    _Investor_AddMoneyInvestor_Handler,
		},
		{
			MethodName: "RemoveMoneyInvestor",
			Handler:    _Investor_RemoveMoneyInvestor_Handler,
		},
		{
			MethodName: "SetMoneyInvestor",
			Handler:    _Investor_SetMoneyInvestor_Handler,
		},
		{
			MethodName: "ListInvestor",
			Handler:    _Investor_ListInvestor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "investor/investor.proto",
}
