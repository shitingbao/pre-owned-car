package service

import (
	"context"
	"log"

	pb "car/api/car"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CarService struct {
	pb.UnimplementedCarServer
}

func NewCarService() *CarService {
	return &CarService{}
}

func (s *CarService) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarReply, error) {
	log.Println("you create car:", req)
	return &pb.CreateCarReply{
		CarId: 123,
	}, nil
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
	log.Println("you car:", req.CreateTime.AsTime())
	return &pb.ListCarReply{
		CarId:      1,
		CarName:    "new car",
		CreateTime: timestamppb.Now(),
		CarModel:   []string{"aa", "bb"},
		CarParam:   map[string]string{"cc": "dd"},
	}, nil
}

func (s *CarService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
