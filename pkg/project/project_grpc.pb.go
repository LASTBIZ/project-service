// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: project/project.proto

package project

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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	GetProjectByID(ctx context.Context, in *GetProjectByIdRequest, opts ...grpc.CallOption) (*GetProjectByIdResponse, error)
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error)
	AddRoadmap(ctx context.Context, in *CreateRoadmapRequest, opts ...grpc.CallOption) (*RoadmapResponse, error)
	RemoveRoadmap(ctx context.Context, in *DeleteRoadmapRequest, opts ...grpc.CallOption) (*RoadmapResponse, error)
	AddInvestor(ctx context.Context, in *AddInvestorRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
	RemoveInvestor(ctx context.Context, in *RemoveInvestorRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
	AddMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
	RemoveMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
	SetMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
	InvestProject(ctx context.Context, in *InvestProjectRequest, opts ...grpc.CallOption) (*InvestorResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjectByID(ctx context.Context, in *GetProjectByIdRequest, opts ...grpc.CallOption) (*GetProjectByIdResponse, error) {
	out := new(GetProjectByIdResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/GetProjectByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error) {
	out := new(GetProjectsResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/GetProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddRoadmap(ctx context.Context, in *CreateRoadmapRequest, opts ...grpc.CallOption) (*RoadmapResponse, error) {
	out := new(RoadmapResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/AddRoadmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) RemoveRoadmap(ctx context.Context, in *DeleteRoadmapRequest, opts ...grpc.CallOption) (*RoadmapResponse, error) {
	out := new(RoadmapResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/RemoveRoadmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddInvestor(ctx context.Context, in *AddInvestorRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/AddInvestor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) RemoveInvestor(ctx context.Context, in *RemoveInvestorRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/RemoveInvestor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) AddMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/AddMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) RemoveMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/RemoveMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SetMoney(ctx context.Context, in *MoneyRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/SetMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) InvestProject(ctx context.Context, in *InvestProjectRequest, opts ...grpc.CallOption) (*InvestorResponse, error) {
	out := new(InvestorResponse)
	err := c.cc.Invoke(ctx, "/project.ProjectService/InvestProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	GetProjectByID(context.Context, *GetProjectByIdRequest) (*GetProjectByIdResponse, error)
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error)
	AddRoadmap(context.Context, *CreateRoadmapRequest) (*RoadmapResponse, error)
	RemoveRoadmap(context.Context, *DeleteRoadmapRequest) (*RoadmapResponse, error)
	AddInvestor(context.Context, *AddInvestorRequest) (*InvestorResponse, error)
	RemoveInvestor(context.Context, *RemoveInvestorRequest) (*InvestorResponse, error)
	AddMoney(context.Context, *MoneyRequest) (*InvestorResponse, error)
	RemoveMoney(context.Context, *MoneyRequest) (*InvestorResponse, error)
	SetMoney(context.Context, *MoneyRequest) (*InvestorResponse, error)
	InvestProject(context.Context, *InvestProjectRequest) (*InvestorResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProjectByID(context.Context, *GetProjectByIdRequest) (*GetProjectByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectByID not implemented")
}
func (UnimplementedProjectServiceServer) GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedProjectServiceServer) AddRoadmap(context.Context, *CreateRoadmapRequest) (*RoadmapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoadmap not implemented")
}
func (UnimplementedProjectServiceServer) RemoveRoadmap(context.Context, *DeleteRoadmapRequest) (*RoadmapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoadmap not implemented")
}
func (UnimplementedProjectServiceServer) AddInvestor(context.Context, *AddInvestorRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInvestor not implemented")
}
func (UnimplementedProjectServiceServer) RemoveInvestor(context.Context, *RemoveInvestorRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInvestor not implemented")
}
func (UnimplementedProjectServiceServer) AddMoney(context.Context, *MoneyRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMoney not implemented")
}
func (UnimplementedProjectServiceServer) RemoveMoney(context.Context, *MoneyRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMoney not implemented")
}
func (UnimplementedProjectServiceServer) SetMoney(context.Context, *MoneyRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMoney not implemented")
}
func (UnimplementedProjectServiceServer) InvestProject(context.Context, *InvestProjectRequest) (*InvestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvestProject not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjectByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjectByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/GetProjectByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjectByID(ctx, req.(*GetProjectByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/GetProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddRoadmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddRoadmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/AddRoadmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddRoadmap(ctx, req.(*CreateRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_RemoveRoadmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoadmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).RemoveRoadmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/RemoveRoadmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).RemoveRoadmap(ctx, req.(*DeleteRoadmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/AddInvestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddInvestor(ctx, req.(*AddInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_RemoveInvestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInvestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).RemoveInvestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/RemoveInvestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).RemoveInvestor(ctx, req.(*RemoveInvestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_AddMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).AddMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/AddMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).AddMoney(ctx, req.(*MoneyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_RemoveMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).RemoveMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/RemoveMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).RemoveMoney(ctx, req.(*MoneyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SetMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoneyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SetMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/SetMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SetMoney(ctx, req.(*MoneyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_InvestProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvestProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).InvestProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/project.ProjectService/InvestProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).InvestProject(ctx, req.(*InvestProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "project.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectService_UpdateProject_Handler,
		},
		{
			MethodName: "GetProjectByID",
			Handler:    _ProjectService_GetProjectByID_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _ProjectService_GetProjects_Handler,
		},
		{
			MethodName: "AddRoadmap",
			Handler:    _ProjectService_AddRoadmap_Handler,
		},
		{
			MethodName: "RemoveRoadmap",
			Handler:    _ProjectService_RemoveRoadmap_Handler,
		},
		{
			MethodName: "AddInvestor",
			Handler:    _ProjectService_AddInvestor_Handler,
		},
		{
			MethodName: "RemoveInvestor",
			Handler:    _ProjectService_RemoveInvestor_Handler,
		},
		{
			MethodName: "AddMoney",
			Handler:    _ProjectService_AddMoney_Handler,
		},
		{
			MethodName: "RemoveMoney",
			Handler:    _ProjectService_RemoveMoney_Handler,
		},
		{
			MethodName: "SetMoney",
			Handler:    _ProjectService_SetMoney_Handler,
		},
		{
			MethodName: "InvestProject",
			Handler:    _ProjectService_InvestProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/project.proto",
}
