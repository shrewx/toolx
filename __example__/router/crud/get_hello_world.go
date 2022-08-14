package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
)

func init() {
	Router.Register(&GetHelloWorld{})
}

type GetHelloWorld struct {
	ginx.MethodGet
}

type GetHelloWorldResponse struct {
	Message string `json:"message"`
}

func (g *GetHelloWorld) Path() string {
	return "/hello"
}

func (g *GetHelloWorld) Output(ctx *gin.Context) (interface{}, error) {
	return GetHelloWorldResponse{Message: "hello world"}, nil
}
