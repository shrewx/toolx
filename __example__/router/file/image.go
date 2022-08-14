package file

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
	"os"
)

func init() {
	Router.Register(&Image{})
}

type Image struct {
	ginx.MethodGet
}

func (g *Image) Path() string {
	return "/image"
}

func (g *Image) Output(ctx *gin.Context) (interface{}, error) {
	png := ginx.NewImagePNG()
	file, err := os.ReadFile("./router/file/go.png")
	if err != nil {
		return nil, err
	}
	png.Write(file)
	return png, nil
}
