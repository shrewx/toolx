package router

import (
	"github.com/shrewx/ginx"
	"github.com/shrewx/toolx/__example__/router/crud"
	"github.com/shrewx/toolx/__example__/router/file"
	"github.com/shrewx/toolx/__example__/router/middleware"
)

var V0Router = ginx.NewRouter(ginx.Group("v0"), &middleware.BaseAuth{})

func init() {
	V0Router.Register(crud.Router)
	V0Router.Register(file.Router)
}
