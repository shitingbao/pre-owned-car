// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: api/car/car.proto

package car

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
	Car_CreateCar_FullMethodName = "/api.car.Car/CreateCar"
	Car_UpdateCar_FullMethodName = "/api.car.Car/UpdateCar"
	Car_DeleteCar_FullMethodName = "/api.car.Car/DeleteCar"
	Car_GetCar_FullMethodName    = "/api.car.Car/GetCar"
	Car_ListCar_FullMethodName   = "/api.car.Car/ListCar"
	Car_SayHello_FullMethodName  = "/api.car.Car/SayHello"
)

// CarClient is the client API for Car service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarClient interface {
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarReply, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarReply, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarReply, error)
	GetCar(ctx context.Context, in *GetCarRequest, opts ...grpc.CallOption) (*GetCarReply, error)
	ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (*ListCarReply, error)
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type carClient struct {
	cc grpc.ClientConnInterface
}

func NewCarClient(cc grpc.ClientConnInterface) CarClient {
	return &carClient{cc}
}

func (c *carClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarReply, error) {
	out := new(CreateCarReply)
	err := c.cc.Invoke(ctx, Car_CreateCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarReply, error) {
	out := new(UpdateCarReply)
	err := c.cc.Invoke(ctx, Car_UpdateCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarReply, error) {
	out := new(DeleteCarReply)
	err := c.cc.Invoke(ctx, Car_DeleteCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) GetCar(ctx context.Context, in *GetCarRequest, opts ...grpc.CallOption) (*GetCarReply, error) {
	out := new(GetCarReply)
	err := c.cc.Invoke(ctx, Car_GetCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (*ListCarReply, error) {
	out := new(ListCarReply)
	err := c.cc.Invoke(ctx, Car_ListCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Car_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServer is the server API for Car service.
// All implementations must embed UnimplementedCarServer
// for forward compatibility
type CarServer interface {
	CreateCar(context.Context, *CreateCarRequest) (*CreateCarReply, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarReply, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarReply, error)
	GetCar(context.Context, *GetCarRequest) (*GetCarReply, error)
	ListCar(context.Context, *ListCarRequest) (*ListCarReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedCarServer()
}

// UnimplementedCarServer must be embedded to have forward compatible implementations.
type UnimplementedCarServer struct {
}

func (UnimplementedCarServer) CreateCar(context.Context, *CreateCarRequest) (*CreateCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedCarServer) UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (UnimplementedCarServer) DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (UnimplementedCarServer) GetCar(context.Context, *GetCarRequest) (*GetCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}
func (UnimplementedCarServer) ListCar(context.Context, *ListCarRequest) (*ListCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCar not implemented")
}
func (UnimplementedCarServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCarServer) mustEmbedUnimplementedCarServer() {}

// UnsafeCarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServer will
// result in compilation errors.
type UnsafeCarServer interface {
	mustEmbedUnimplementedCarServer()
}

func RegisterCarServer(s grpc.ServiceRegistrar, srv CarServer) {
	s.RegisterService(&Car_ServiceDesc, srv)
}

func _Car_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_CreateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_UpdateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_DeleteCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_GetCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).GetCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_GetCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).GetCar(ctx, req.(*GetCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_ListCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).ListCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_ListCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).ListCar(ctx, req.(*ListCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Car_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Car_ServiceDesc is the grpc.ServiceDesc for Car service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Car_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.car.Car",
	HandlerType: (*CarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCar",
			Handler:    _Car_CreateCar_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _Car_UpdateCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _Car_DeleteCar_Handler,
		},
		{
			MethodName: "GetCar",
			Handler:    _Car_GetCar_Handler,
		},
		{
			MethodName: "ListCar",
			Handler:    _Car_ListCar_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Car_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/car/car.proto",
}
