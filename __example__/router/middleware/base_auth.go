package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/shrewx/ginx"
	"github.com/shrewx/ginx/pkg/errors"
	"unsafe"
)

type BaseAuth struct {
	ginx.HTTPBasicAuthSecurityType
	Name string `in:"header" name:"Authorization"`
}

func (g *BaseAuth) Output(ctx *gin.Context) (interface{}, error) {
	if g.Name != authorization("admin", "admin") {
		return nil, errors.Unauthorized
	}

	return nil, nil
}

func authorization(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString(stringToBytes(base))
}

func stringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
