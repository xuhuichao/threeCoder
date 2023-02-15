// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: likeService.proto

package service

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

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	CreateLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error)
	DeleteLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error)
	FindVideoIds(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeDetailResponse, error)
	IsLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) CreateLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error) {
	out := new(LikeCommonResponse)
	err := c.cc.Invoke(ctx, "/pb.LikeService/CreateLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) DeleteLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error) {
	out := new(LikeCommonResponse)
	err := c.cc.Invoke(ctx, "/pb.LikeService/DeleteLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) FindVideoIds(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeDetailResponse, error) {
	out := new(LikeDetailResponse)
	err := c.cc.Invoke(ctx, "/pb.LikeService/FindVideoIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) IsLike(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeCommonResponse, error) {
	out := new(LikeCommonResponse)
	err := c.cc.Invoke(ctx, "/pb.LikeService/IsLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility
type LikeServiceServer interface {
	CreateLike(context.Context, *LikeRequest) (*LikeCommonResponse, error)
	DeleteLike(context.Context, *LikeRequest) (*LikeCommonResponse, error)
	FindVideoIds(context.Context, *LikeRequest) (*LikeDetailResponse, error)
	IsLike(context.Context, *LikeRequest) (*LikeCommonResponse, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServiceServer struct {
}

func (UnimplementedLikeServiceServer) CreateLike(context.Context, *LikeRequest) (*LikeCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLike not implemented")
}
func (UnimplementedLikeServiceServer) DeleteLike(context.Context, *LikeRequest) (*LikeCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLike not implemented")
}
func (UnimplementedLikeServiceServer) FindVideoIds(context.Context, *LikeRequest) (*LikeDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindVideoIds not implemented")
}
func (UnimplementedLikeServiceServer) IsLike(context.Context, *LikeRequest) (*LikeCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLike not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_CreateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).CreateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LikeService/CreateLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).CreateLike(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_DeleteLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).DeleteLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LikeService/DeleteLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).DeleteLike(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_FindVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).FindVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LikeService/FindVideoIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).FindVideoIds(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_IsLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).IsLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LikeService/IsLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).IsLike(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLike",
			Handler:    _LikeService_CreateLike_Handler,
		},
		{
			MethodName: "DeleteLike",
			Handler:    _LikeService_DeleteLike_Handler,
		},
		{
			MethodName: "FindVideoIds",
			Handler:    _LikeService_FindVideoIds_Handler,
		},
		{
			MethodName: "IsLike",
			Handler:    _LikeService_IsLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "likeService.proto",
}