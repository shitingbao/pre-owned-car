package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// UploadRepo is a Greater repo.
type UploadRepo interface {
	FormFile(ctx http.Context) error
}

// UploadUsecase is a Car usecase.
type UploadUsecase struct {
	repo UploadRepo
	log  *log.Helper
}

// NewUploadUsecase new a upload usecase.
func NewUploadUsecase(repo UploadRepo, logger log.Logger) *UploadUsecase {
	return &UploadUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Upload, form file.
func (uc *UploadUsecase) FormFile(ctx http.Context) error {
	// uc.log.WithContext(ctx).Infof("C: %v", g.Name)

	return uc.repo.FormFile(ctx)
}
