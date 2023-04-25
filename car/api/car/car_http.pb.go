// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v4.22.3
// source: api/car/car.proto

package car

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCarCreateCar = "/api.car.Car/CreateCar"
const OperationCarListCar = "/api.car.Car/ListCar"
const OperationCarSayHello = "/api.car.Car/SayHello"

type CarHTTPServer interface {
	CreateCar(context.Context, *CreateCarRequest) (*CreateCarReply, error)
	ListCar(context.Context, *ListCarRequest) (*ListCarReply, error)
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterCarHTTPServer(s *http.Server, srv CarHTTPServer) {
	r := s.Route("/")
	r.POST("/car", _Car_CreateCar0_HTTP_Handler(srv))
	r.GET("/car/list", _Car_ListCar0_HTTP_Handler(srv))
	r.GET("/car/helloworld/{name}", _Car_SayHello0_HTTP_Handler(srv))
}

func _Car_CreateCar0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCarRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarCreateCar)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCar(ctx, req.(*CreateCarRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateCarReply)
		return ctx.Result(200, reply)
	}
}

func _Car_ListCar0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCarRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarListCar)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCar(ctx, req.(*ListCarRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCarReply)
		return ctx.Result(200, reply)
	}
}

func _Car_SayHello0_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarSayHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type CarHTTPClient interface {
	CreateCar(ctx context.Context, req *CreateCarRequest, opts ...http.CallOption) (rsp *CreateCarReply, err error)
	ListCar(ctx context.Context, req *ListCarRequest, opts ...http.CallOption) (rsp *ListCarReply, err error)
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type CarHTTPClientImpl struct {
	cc *http.Client
}

func NewCarHTTPClient(client *http.Client) CarHTTPClient {
	return &CarHTTPClientImpl{client}
}

func (c *CarHTTPClientImpl) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...http.CallOption) (*CreateCarReply, error) {
	var out CreateCarReply
	pattern := "/car"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCarCreateCar))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) ListCar(ctx context.Context, in *ListCarRequest, opts ...http.CallOption) (*ListCarReply, error) {
	var out ListCarReply
	pattern := "/car/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarListCar))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/car/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarSayHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
