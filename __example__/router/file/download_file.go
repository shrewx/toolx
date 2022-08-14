package file

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
)

func init() {
	Router.Register(&DownloadFile{})
}

type DownloadFile struct {
	ginx.MethodGet
}

func (g *DownloadFile) Path() string {
	return ""
}

func (g *DownloadFile) Output(ctx *gin.Context) (interface{}, error) {
	file := ginx.NewAttachment("text.txt", ginx.MineApplicationOctetStream)
	file.Write([]byte("hello world"))
	return file, nil
}
