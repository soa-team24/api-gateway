// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: stakeholder/stakeholder-service.proto

package stakeholder

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
	StakeholderService_RegisterTourist_FullMethodName = "/StakeholderService/RegisterTourist"
	StakeholderService_Login_FullMethodName           = "/StakeholderService/Login"
)

// StakeholderServiceClient is the client API for StakeholderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StakeholderServiceClient interface {
	RegisterTourist(ctx context.Context, in *RegisterTouristRequest, opts ...grpc.CallOption) (*RegisterTouristResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*RegisterTouristResponse, error)
}

type stakeholderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStakeholderServiceClient(cc grpc.ClientConnInterface) StakeholderServiceClient {
	return &stakeholderServiceClient{cc}
}

func (c *stakeholderServiceClient) RegisterTourist(ctx context.Context, in *RegisterTouristRequest, opts ...grpc.CallOption) (*RegisterTouristResponse, error) {
	out := new(RegisterTouristResponse)
	err := c.cc.Invoke(ctx, StakeholderService_RegisterTourist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakeholderServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*RegisterTouristResponse, error) {
	out := new(RegisterTouristResponse)
	err := c.cc.Invoke(ctx, StakeholderService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakeholderServiceServer is the server API for StakeholderService service.
// All implementations must embed UnimplementedStakeholderServiceServer
// for forward compatibility
type StakeholderServiceServer interface {
	RegisterTourist(context.Context, *RegisterTouristRequest) (*RegisterTouristResponse, error)
	Login(context.Context, *LoginRequest) (*RegisterTouristResponse, error)
	mustEmbedUnimplementedStakeholderServiceServer()
}

// UnimplementedStakeholderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStakeholderServiceServer struct {
}

func (UnimplementedStakeholderServiceServer) RegisterTourist(context.Context, *RegisterTouristRequest) (*RegisterTouristResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTourist not implemented")
}
func (UnimplementedStakeholderServiceServer) Login(context.Context, *LoginRequest) (*RegisterTouristResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedStakeholderServiceServer) mustEmbedUnimplementedStakeholderServiceServer() {}

// UnsafeStakeholderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StakeholderServiceServer will
// result in compilation errors.
type UnsafeStakeholderServiceServer interface {
	mustEmbedUnimplementedStakeholderServiceServer()
}

func RegisterStakeholderServiceServer(s grpc.ServiceRegistrar, srv StakeholderServiceServer) {
	s.RegisterService(&StakeholderService_ServiceDesc, srv)
}

func _StakeholderService_RegisterTourist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTouristRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).RegisterTourist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_RegisterTourist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).RegisterTourist(ctx, req.(*RegisterTouristRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakeholderService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StakeholderService_ServiceDesc is the grpc.ServiceDesc for StakeholderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StakeholderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StakeholderService",
	HandlerType: (*StakeholderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterTourist",
			Handler:    _StakeholderService_RegisterTourist_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _StakeholderService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stakeholder/stakeholder-service.proto",
}
