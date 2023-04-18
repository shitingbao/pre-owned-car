package service

import (
	"context"

	pb "car/api/car"
)

type CarService struct {
	pb.UnimplementedCarServer
}

func NewCarService() *CarService {
	return &CarService{}
}

func (s *CarService) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarReply, error) {
	return &pb.CreateCarReply{}, nil
}

func (s *CarService) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.UpdateCarReply, error) {
	return &pb.UpdateCarReply{}, nil
}

func (s *CarService) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteCarReply, error) {
	return &pb.DeleteCarReply{}, nil
}

func (s *CarService) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.GetCarReply, error) {
	return &pb.GetCarReply{}, nil
}

func (s *CarService) ListCar(ctx context.Context, req *pb.ListCarRequest) (*pb.ListCarReply, error) {
	return &pb.ListCarReply{}, nil
}

func (s *CarService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
