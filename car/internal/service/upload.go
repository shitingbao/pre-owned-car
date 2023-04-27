package service

import (
	"car/internal/biz"

	"github.com/go-kratos/kratos/v2/transport/http"
)

type UploadService struct {
	up *biz.UploadUsecase
}

func NewUploadService(c *biz.UploadUsecase) *UploadService {
	return &UploadService{
		up: c,
	}
}

func (s *UploadService) FormFile(ctx http.Context) error {

	return s.up.FormFile(ctx)
}
