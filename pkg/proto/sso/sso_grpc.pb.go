// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/sso/sso.proto

package sso

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
	SSOService_CheckPermission_FullMethodName = "/sso.v1.SSOService/CheckPermission"
	SSOService_GetUserByUID_FullMethodName    = "/sso.v1.SSOService/GetUserByUID"
	SSOService_GetRolesByUID_FullMethodName   = "/sso.v1.SSOService/GetRolesByUID"
	SSOService_GetUsers_FullMethodName        = "/sso.v1.SSOService/GetUsers"
	SSOService_GetGroupsDetail_FullMethodName = "/sso.v1.SSOService/GetGroupsDetail"
)

// SSOServiceClient is the client API for SSOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SSOServiceClient interface {
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
	GetUserByUID(ctx context.Context, in *GetUserByUIDRequest, opts ...grpc.CallOption) (*GetUserByUIDResponse, error)
	GetRolesByUID(ctx context.Context, in *GetRolesByUIDRequest, opts ...grpc.CallOption) (*GetRolesByUIDResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetGroupsDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsDetailResponse, error)
}

type sSOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSOServiceClient(cc grpc.ClientConnInterface) SSOServiceClient {
	return &sSOServiceClient{cc}
}

func (c *sSOServiceClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, SSOService_CheckPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) GetUserByUID(ctx context.Context, in *GetUserByUIDRequest, opts ...grpc.CallOption) (*GetUserByUIDResponse, error) {
	out := new(GetUserByUIDResponse)
	err := c.cc.Invoke(ctx, SSOService_GetUserByUID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) GetRolesByUID(ctx context.Context, in *GetRolesByUIDRequest, opts ...grpc.CallOption) (*GetRolesByUIDResponse, error) {
	out := new(GetRolesByUIDResponse)
	err := c.cc.Invoke(ctx, SSOService_GetRolesByUID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, SSOService_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) GetGroupsDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsDetailResponse, error) {
	out := new(GetGroupsDetailResponse)
	err := c.cc.Invoke(ctx, SSOService_GetGroupsDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSOServiceServer is the server API for SSOService service.
// All implementations must embed UnimplementedSSOServiceServer
// for forward compatibility
type SSOServiceServer interface {
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	GetUserByUID(context.Context, *GetUserByUIDRequest) (*GetUserByUIDResponse, error)
	GetRolesByUID(context.Context, *GetRolesByUIDRequest) (*GetRolesByUIDResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetGroupsDetail(context.Context, *emptypb.Empty) (*GetGroupsDetailResponse, error)
	mustEmbedUnimplementedSSOServiceServer()
}

// UnimplementedSSOServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSSOServiceServer struct {
}

func (UnimplementedSSOServiceServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedSSOServiceServer) GetUserByUID(context.Context, *GetUserByUIDRequest) (*GetUserByUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUID not implemented")
}
func (UnimplementedSSOServiceServer) GetRolesByUID(context.Context, *GetRolesByUIDRequest) (*GetRolesByUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesByUID not implemented")
}
func (UnimplementedSSOServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedSSOServiceServer) GetGroupsDetail(context.Context, *emptypb.Empty) (*GetGroupsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupsDetail not implemented")
}
func (UnimplementedSSOServiceServer) mustEmbedUnimplementedSSOServiceServer() {}

// UnsafeSSOServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSOServiceServer will
// result in compilation errors.
type UnsafeSSOServiceServer interface {
	mustEmbedUnimplementedSSOServiceServer()
}

func RegisterSSOServiceServer(s grpc.ServiceRegistrar, srv SSOServiceServer) {
	s.RegisterService(&SSOService_ServiceDesc, srv)
}

func _SSOService_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_CheckPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_GetUserByUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).GetUserByUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_GetUserByUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).GetUserByUID(ctx, req.(*GetUserByUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_GetRolesByUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesByUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).GetRolesByUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_GetRolesByUID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).GetRolesByUID(ctx, req.(*GetRolesByUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_GetGroupsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).GetGroupsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_GetGroupsDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).GetGroupsDetail(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SSOService_ServiceDesc is the grpc.ServiceDesc for SSOService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSOService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sso.v1.SSOService",
	HandlerType: (*SSOServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPermission",
			Handler:    _SSOService_CheckPermission_Handler,
		},
		{
			MethodName: "GetUserByUID",
			Handler:    _SSOService_GetUserByUID_Handler,
		},
		{
			MethodName: "GetRolesByUID",
			Handler:    _SSOService_GetRolesByUID_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _SSOService_GetUsers_Handler,
		},
		{
			MethodName: "GetGroupsDetail",
			Handler:    _SSOService_GetGroupsDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sso/sso.proto",
}
