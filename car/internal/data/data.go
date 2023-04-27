package data

import (
	"car/internal/conf"
	"log"

	glog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCarRepo, NewUploadRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger glog.Logger) (*Data, func(), error) {
	log.Println("c.Database.Source:", c.Database.Source)
	d, err := OpenMysqlClient(c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		glog.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: d,
	}, cleanup, nil
}

func OpenMysqlClient(source string) (*gorm.DB, error) {
	d, err := gorm.Open(mysql.Open(source), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	d.Debug()
	db, err := d.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	return d, nil
}
