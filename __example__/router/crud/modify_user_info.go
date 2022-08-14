package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
)

func init() {
	Router.Register(&ModifyUserInfo{})
}

type ModifyUserInfo struct {
	ginx.MethodPut
	Name string `in:"urlencoded" name:"name"`
}

func (g *ModifyUserInfo) Path() string {
	return "/:id"
}

func (g *ModifyUserInfo) Output(ctx *gin.Context) (interface{}, error) {
	return g.Name, nil
}
