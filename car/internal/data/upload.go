package data

import (
	"car/internal/biz"
	"io"
	"os"

	glog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type uploadRepo struct {
	log *glog.Helper
}

// NewUploadRepo .
func NewUploadRepo(logger glog.Logger) biz.UploadRepo {
	return &uploadRepo{
		log: glog.NewHelper(logger),
	}
}

// FormFile 表单文件
// example post 127.0.0.1:8000/upload
// form-data
// keytype txt key name value test
// keytype file key file value file
func (r *uploadRepo) FormFile(ctx http.Context) error {
	req := ctx.Request()

	fileName := req.FormValue("name")
	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}
	return ctx.JSON(200, map[string]string{"result": "File " + fileName + " Uploaded successfully"})
}
