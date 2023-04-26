package service

import (
	"context"
	"log"

	pb "car/api/car"
	"car/internal/biz"
)

type CarService struct {
	pb.UnimplementedCarServer

	car *biz.CarUsecase
}

func NewCarService(c *biz.CarUsecase) *CarService {
	return &CarService{
		car: c,
	}
}

func (s *CarService) CreateCar(ctx context.Context, req *pb.CreateCarRequest) (*pb.CreateCarReply, error) {
	c, err := s.car.CreateCar(ctx, &biz.Car{
		Name: req.CarName,
	})
	return &pb.CreateCarReply{
		CarId: c.ID,
	}, err
}

func (s *CarService) UpdateCar(ctx context.Context, req *pb.UpdateCarRequest) (*pb.UpdateCarReply, error) {
	c, err := s.car.Update(ctx, &biz.Car{
		ID:   req.Id,
		Name: req.Name,
	})
	return &pb.UpdateCarReply{
		Name: c.Name,
	}, err
}

func (s *CarService) DeleteCar(ctx context.Context, req *pb.DeleteCarRequest) (*pb.DeleteCarReply, error) {
	return &pb.DeleteCarReply{}, nil
}

func (s *CarService) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.GetCarReply, error) {
	c, err := s.car.FindByID(ctx, req.Id)
	return &pb.GetCarReply{
		Name: c.Name,
	}, err
}

func (s *CarService) ListCar(ctx context.Context, req *pb.ListCarRequest) (*pb.ListCarReply, error) {
	log.Println("you car:", req.CreateTime.AsTime())
	res, err := s.car.ListAll(ctx)
	list := []*pb.CarReply{}
	for _, v := range res {
		p := &pb.CarReply{
			CarName: v.Name,
		}
		list = append(list, p)
	}
	return &pb.ListCarReply{
		List: list,
	}, err
}

func (s *CarService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{}, nil
}
