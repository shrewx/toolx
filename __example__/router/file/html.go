package file

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
)

func init() {
	Router.Register(&HTML{})
}

type HTML struct {
	ginx.MethodGet
}

func (g *HTML) Path() string {
	return "/html"
}

func (g *HTML) Output(ctx *gin.Context) (interface{}, error) {
	file := ginx.NewHTML()
	file.Write([]byte("hello world"))
	return file, nil
}
