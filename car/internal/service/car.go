package service

import (
	"context"
	"log"

	api_car "car/api/car"
	"car/internal/biz"
)

type CarService struct {
	api_car.UnimplementedCarServer

	car *biz.CarUsecase
}

func NewCarService(c *biz.CarUsecase) *CarService {
	return &CarService{
		car: c,
	}
}

func (s *CarService) CreateCar(ctx context.Context, req *api_car.CreateCarRequest) (*api_car.CreateCarReply, error) {
	c, err := s.car.CreateCar(ctx, &biz.Car{
		Name: req.CarName,
	})
	return &api_car.CreateCarReply{
		CarId: c.ID,
	}, err
}

func (s *CarService) UpdateCar(ctx context.Context, req *api_car.UpdateCarRequest) (*api_car.UpdateCarReply, error) {
	c, err := s.car.Update(ctx, &biz.Car{
		ID:   req.Id,
		Name: req.Name,
	})
	return &api_car.UpdateCarReply{
		Name: c.Name,
	}, err
}

func (s *CarService) DeleteCar(ctx context.Context, req *api_car.DeleteCarRequest) (*api_car.DeleteCarReply, error) {
	return &api_car.DeleteCarReply{}, nil
}

func (s *CarService) GetCar(ctx context.Context, req *api_car.GetCarRequest) (*api_car.GetCarReply, error) {
	// if req.Id == 0 {
	// 	return nil, api_car.ErrorUserNotFound("not id", "no this car")
	// }
	cars, err := s.car.Find(ctx, req)
	return &api_car.GetCarReply{
		Cars: cars,
	}, err
}

func (s *CarService) ListCar(ctx context.Context, req *api_car.ListCarRequest) (*api_car.ListCarReply, error) {
	log.Println("you car:", req.CreateTime.AsTime())
	res, err := s.car.ListAll(ctx)
	list := []*api_car.CarReply{}
	for _, v := range res {
		p := &api_car.CarReply{
			CarName: v.Name,
		}
		list = append(list, p)
	}
	return &api_car.ListCarReply{
		List: list,
	}, err
}

func (s *CarService) SayHello(ctx context.Context, req *api_car.HelloRequest) (*api_car.HelloReply, error) {
	return &api_car.HelloReply{}, nil
}
