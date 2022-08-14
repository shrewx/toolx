package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
)

func init() {
	Router.Register(&GetUserInfo{})
}

type GetUserInfo struct {
	ginx.MethodGet
	Username string `in:"query" name:"username"`
	ID       int    `in:"path" name:"id"`
}

type GetUserInfoResponse struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
}

func (g *GetUserInfo) Path() string {
	return "/:id"
}

func (g *GetUserInfo) Output(ctx *gin.Context) (interface{}, error) {
	return GetUserInfoResponse{
		Username: g.Username,
		ID:       g.ID,
	}, nil
}
