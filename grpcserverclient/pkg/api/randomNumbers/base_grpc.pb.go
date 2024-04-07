// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: base.proto

package randomNumbers

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
	RandomNumbers_Generate_FullMethodName        = "/api.RandomNumbers/Generate"
	RandomNumbers_GenerateCrypto_FullMethodName  = "/api.RandomNumbers/GenerateCrypto"
	RandomNumbers_GenerateNumbers_FullMethodName = "/api.RandomNumbers/GenerateNumbers"
	RandomNumbers_SendNumbers_FullMethodName     = "/api.RandomNumbers/SendNumbers"
)

// RandomNumbersClient is the client API for RandomNumbers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomNumbersClient interface {
	Generate(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (*GenResponse, error)
	GenerateCrypto(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error)
	GenerateNumbers(ctx context.Context, in *GenNumRequest, opts ...grpc.CallOption) (*GenNumResponse, error)
	SendNumbers(ctx context.Context, in *SendNumRequest, opts ...grpc.CallOption) (*SendNumResponse, error)
}

type randomNumbersClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomNumbersClient(cc grpc.ClientConnInterface) RandomNumbersClient {
	return &randomNumbersClient{cc}
}

func (c *randomNumbersClient) Generate(ctx context.Context, in *GenRequest, opts ...grpc.CallOption) (*GenResponse, error) {
	out := new(GenResponse)
	err := c.cc.Invoke(ctx, RandomNumbers_Generate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomNumbersClient) GenerateCrypto(ctx context.Context, in *CryptRequest, opts ...grpc.CallOption) (*CryptResponse, error) {
	out := new(CryptResponse)
	err := c.cc.Invoke(ctx, RandomNumbers_GenerateCrypto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomNumbersClient) GenerateNumbers(ctx context.Context, in *GenNumRequest, opts ...grpc.CallOption) (*GenNumResponse, error) {
	out := new(GenNumResponse)
	err := c.cc.Invoke(ctx, RandomNumbers_GenerateNumbers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *randomNumbersClient) SendNumbers(ctx context.Context, in *SendNumRequest, opts ...grpc.CallOption) (*SendNumResponse, error) {
	out := new(SendNumResponse)
	err := c.cc.Invoke(ctx, RandomNumbers_SendNumbers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomNumbersServer is the server API for RandomNumbers service.
// All implementations must embed UnimplementedRandomNumbersServer
// for forward compatibility
type RandomNumbersServer interface {
	Generate(context.Context, *GenRequest) (*GenResponse, error)
	GenerateCrypto(context.Context, *CryptRequest) (*CryptResponse, error)
	GenerateNumbers(context.Context, *GenNumRequest) (*GenNumResponse, error)
	SendNumbers(context.Context, *SendNumRequest) (*SendNumResponse, error)
	mustEmbedUnimplementedRandomNumbersServer()
}

// UnimplementedRandomNumbersServer must be embedded to have forward compatible implementations.
type UnimplementedRandomNumbersServer struct {
}

func (UnimplementedRandomNumbersServer) Generate(context.Context, *GenRequest) (*GenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedRandomNumbersServer) GenerateCrypto(context.Context, *CryptRequest) (*CryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCrypto not implemented")
}
func (UnimplementedRandomNumbersServer) GenerateNumbers(context.Context, *GenNumRequest) (*GenNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateNumbers not implemented")
}
func (UnimplementedRandomNumbersServer) SendNumbers(context.Context, *SendNumRequest) (*SendNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNumbers not implemented")
}
func (UnimplementedRandomNumbersServer) mustEmbedUnimplementedRandomNumbersServer() {}

// UnsafeRandomNumbersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomNumbersServer will
// result in compilation errors.
type UnsafeRandomNumbersServer interface {
	mustEmbedUnimplementedRandomNumbersServer()
}

func RegisterRandomNumbersServer(s grpc.ServiceRegistrar, srv RandomNumbersServer) {
	s.RegisterService(&RandomNumbers_ServiceDesc, srv)
}

func _RandomNumbers_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumbersServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomNumbers_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumbersServer).Generate(ctx, req.(*GenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomNumbers_GenerateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumbersServer).GenerateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomNumbers_GenerateCrypto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumbersServer).GenerateCrypto(ctx, req.(*CryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomNumbers_GenerateNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumbersServer).GenerateNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomNumbers_GenerateNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumbersServer).GenerateNumbers(ctx, req.(*GenNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RandomNumbers_SendNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumbersServer).SendNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RandomNumbers_SendNumbers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumbersServer).SendNumbers(ctx, req.(*SendNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RandomNumbers_ServiceDesc is the grpc.ServiceDesc for RandomNumbers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomNumbers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RandomNumbers",
	HandlerType: (*RandomNumbersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _RandomNumbers_Generate_Handler,
		},
		{
			MethodName: "GenerateCrypto",
			Handler:    _RandomNumbers_GenerateCrypto_Handler,
		},
		{
			MethodName: "GenerateNumbers",
			Handler:    _RandomNumbers_GenerateNumbers_Handler,
		},
		{
			MethodName: "SendNumbers",
			Handler:    _RandomNumbers_SendNumbers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base.proto",
}
