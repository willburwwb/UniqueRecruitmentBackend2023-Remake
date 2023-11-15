// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/open_platform/open.proto

package open

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

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSServiceClient interface {
	PushSMS(ctx context.Context, in *PushSMSRequest, opts ...grpc.CallOption) (*PushSMSResponse, error)
	GetAllSMSTemplates(ctx context.Context, in *GetAllSMSTemplatesRequest, opts ...grpc.CallOption) (*GetAllSMSTemplatesResponse, error)
}

type sMSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSServiceClient(cc grpc.ClientConnInterface) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) PushSMS(ctx context.Context, in *PushSMSRequest, opts ...grpc.CallOption) (*PushSMSResponse, error) {
	out := new(PushSMSResponse)
	err := c.cc.Invoke(ctx, "/open_platform.v1.SMSService/PushSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sMSServiceClient) GetAllSMSTemplates(ctx context.Context, in *GetAllSMSTemplatesRequest, opts ...grpc.CallOption) (*GetAllSMSTemplatesResponse, error) {
	out := new(GetAllSMSTemplatesResponse)
	err := c.cc.Invoke(ctx, "/open_platform.v1.SMSService/GetAllSMSTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServiceServer is the server API for SMSService service.
// All implementations must embed UnimplementedSMSServiceServer
// for forward compatibility
type SMSServiceServer interface {
	PushSMS(context.Context, *PushSMSRequest) (*PushSMSResponse, error)
	GetAllSMSTemplates(context.Context, *GetAllSMSTemplatesRequest) (*GetAllSMSTemplatesResponse, error)
	mustEmbedUnimplementedSMSServiceServer()
}

// UnimplementedSMSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (UnimplementedSMSServiceServer) PushSMS(context.Context, *PushSMSRequest) (*PushSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushSMS not implemented")
}
func (UnimplementedSMSServiceServer) GetAllSMSTemplates(context.Context, *GetAllSMSTemplatesRequest) (*GetAllSMSTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSMSTemplates not implemented")
}
func (UnimplementedSMSServiceServer) mustEmbedUnimplementedSMSServiceServer() {}

// UnsafeSMSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSServiceServer will
// result in compilation errors.
type UnsafeSMSServiceServer interface {
	mustEmbedUnimplementedSMSServiceServer()
}

func RegisterSMSServiceServer(s grpc.ServiceRegistrar, srv SMSServiceServer) {
	s.RegisterService(&SMSService_ServiceDesc, srv)
}

func _SMSService_PushSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).PushSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_platform.v1.SMSService/PushSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).PushSMS(ctx, req.(*PushSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SMSService_GetAllSMSTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSMSTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).GetAllSMSTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_platform.v1.SMSService/GetAllSMSTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).GetAllSMSTemplates(ctx, req.(*GetAllSMSTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SMSService_ServiceDesc is the grpc.ServiceDesc for SMSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "open_platform.v1.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushSMS",
			Handler:    _SMSService_PushSMS_Handler,
		},
		{
			MethodName: "GetAllSMSTemplates",
			Handler:    _SMSService_GetAllSMSTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/open_platform/open.proto",
}

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	PushEmail(ctx context.Context, in *PushEmailRequest, opts ...grpc.CallOption) (*PushEmailResponse, error)
	GetAllEmailTemplates(ctx context.Context, in *GetAllEmailTemplatesRequest, opts ...grpc.CallOption) (*GetAllEmailTemplatesResponse, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) PushEmail(ctx context.Context, in *PushEmailRequest, opts ...grpc.CallOption) (*PushEmailResponse, error) {
	out := new(PushEmailResponse)
	err := c.cc.Invoke(ctx, "/open_platform.v1.EmailService/PushEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetAllEmailTemplates(ctx context.Context, in *GetAllEmailTemplatesRequest, opts ...grpc.CallOption) (*GetAllEmailTemplatesResponse, error) {
	out := new(GetAllEmailTemplatesResponse)
	err := c.cc.Invoke(ctx, "/open_platform.v1.EmailService/GetAllEmailTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	PushEmail(context.Context, *PushEmailRequest) (*PushEmailResponse, error)
	GetAllEmailTemplates(context.Context, *GetAllEmailTemplatesRequest) (*GetAllEmailTemplatesResponse, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) PushEmail(context.Context, *PushEmailRequest) (*PushEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushEmail not implemented")
}
func (UnimplementedEmailServiceServer) GetAllEmailTemplates(context.Context, *GetAllEmailTemplatesRequest) (*GetAllEmailTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEmailTemplates not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_PushEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).PushEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_platform.v1.EmailService/PushEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).PushEmail(ctx, req.(*PushEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetAllEmailTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEmailTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetAllEmailTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/open_platform.v1.EmailService/GetAllEmailTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetAllEmailTemplates(ctx, req.(*GetAllEmailTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "open_platform.v1.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushEmail",
			Handler:    _EmailService_PushEmail_Handler,
		},
		{
			MethodName: "GetAllEmailTemplates",
			Handler:    _EmailService_GetAllEmailTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/open_platform/open.proto",
}
