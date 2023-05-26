package server

import (
	carv1 "car/api/car"
	v1 "car/api/helloworld/v1"
	"car/internal/conf"
	"car/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, car *service.CarService, up *service.UploadService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			ratelimit.Server(), // https://go-kratos.dev/docs/component/middleware/ratelimit
			recovery.Recovery(
				recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
					// 异常退出逻辑
					return nil
				}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	route := srv.Route("/")
	route.POST("/upload", up.FormFile)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	carv1.RegisterCarHTTPServer(srv, car)
	return srv
}
