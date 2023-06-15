package data

import (
	api_car "car/api/car"
	"car/internal/biz"
	"context"
	"log"

	glog "github.com/go-kratos/kratos/v2/log"
)

type carRepo struct {
	data *Data
	log  *glog.Helper
}

// NewCarRepo .
func NewCarRepo(data *Data, logger glog.Logger) biz.CarRepo {
	return &carRepo{
		data: data,
		log:  glog.NewHelper(logger),
	}
}

func (r *carRepo) Save(ctx context.Context, g *biz.Car) (*biz.Car, error) {
	err := r.data.db.Create(g).Error
	return g, err
}

func (r *carRepo) Update(ctx context.Context, g *biz.Car) (*biz.Car, error) {
	log.Println(g)
	err := r.data.db.Table("car").Where("id = ?", g.ID).Updates(g).Error
	return g, err
}

func (c *carRepo) Find(ctx context.Context, req *api_car.GetCarRequest) ([]*api_car.CarModel, error) {
	res := []*api_car.CarModel{}
	if err := c.data.db.Table("car").
		Find(req, "price between 1 and 200").
		Scan(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (r *carRepo) FindByID(ctx context.Context, id int64) (*biz.Car, error) {
	c := &biz.Car{}
	err := r.data.db.Table("car").Where("id = ?", id).Scan(c).Error
	return c, err
}

func (r *carRepo) ListAll(context.Context) ([]*biz.Car, error) {
	res := []*biz.Car{}
	err := r.data.db.Table("car").Scan(&res).Error
	return res, err
}
