package crud

import (
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
	"github.com/shrewx/ginx/pkg/errors"
	"github.com/shrewx/toolx/__example__/constant"
)

func init() {
	Router.Register(&CreateUserInfo{})
}

// 创建用户信息
type CreateUserInfo struct {
	ginx.MethodPost
	// Body
	Body struct {
		// 名称
		Name string `json:"name" validate:"required"`
		// 年龄
		Age int `json:"age" validate:"required"`
		// 地址
		Address string `json:"address"`
		// 职业
		Job constant.Job `json:"job" validate:"required"`
		// 城市
		City constant.City `json:"city"`
	} `in:"body"`
}

type CreateUserInfoResponse struct {
	// 用户ID
	ID int `json:"ID"`
	// 名称
	Name string `json:"name"`
	// 年龄
	Age int `json:"age"`
	// 地址
	Address string `json:"address"`
}

func (g *CreateUserInfo) Path() string {
	return ""
}

func (g *CreateUserInfo) Output(ctx *gin.Context) (interface{}, error) {
	if g.Body.Name == "" {
		return nil, errors.BadRequest
	}

	return CreateUserInfoResponse{
		ID:      1,
		Name:    g.Body.Name,
		Age:     g.Body.Age,
		Address: g.Body.Address,
	}, nil
}
