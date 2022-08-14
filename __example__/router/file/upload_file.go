package file

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
	"github.com/shrewx/toolx/__example__/router/errors"
	"mime/multipart"
)

func init() {
	Router.Register(&UploadFile{})
}

type UploadFile struct {
	ginx.MethodPost
	File *multipart.FileHeader `in:"form" name:"file1"`
}

func (u *UploadFile) Path() string {
	return ""
}

func (u *UploadFile) Output(ctx *gin.Context) (interface{}, error) {
	if u.File == nil {
		return nil, errors.UploadFileIsNotExist
	}

	if err := ctx.SaveUploadedFile(u.File, u.File.Filename); err != nil {
		return nil, err
	}

	return nil, nil
}
