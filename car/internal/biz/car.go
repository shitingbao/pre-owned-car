package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Car is a Car model.
type Car struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (Car) TableName() string {
	return "car"
}

// CarRepo is a Greater repo.
type CarRepo interface {
	Save(context.Context, *Car) (*Car, error)
	Update(context.Context, *Car) (*Car, error)
	FindByID(context.Context, int64) (*Car, error)
	ListAll(context.Context) ([]*Car, error)
}

// CarUsecase is a Car usecase.
type CarUsecase struct {
	repo CarRepo
	log  *log.Helper
}

// NewCarUsecase new a Car usecase.
func NewCarUsecase(repo CarRepo, logger log.Logger) *CarUsecase {
	return &CarUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateCar creates a Car, and returns the new Car.
func (uc *CarUsecase) CreateCar(ctx context.Context, g *Car) (*Car, error) {
	uc.log.WithContext(ctx).Infof("CreateCar: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *CarUsecase) Update(ctx context.Context, c *Car) (*Car, error) {
	return uc.repo.Update(ctx, c)
}

func (uc *CarUsecase) FindByID(ctx context.Context, id int64) (*Car, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *CarUsecase) ListAll(ctx context.Context) ([]*Car, error) {
	return uc.repo.ListAll(ctx)
}
