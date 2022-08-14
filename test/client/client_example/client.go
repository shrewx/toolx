package client_example

import (
	context "context"

	github_com_shrewx_ginx "github.com/shrewx/ginx"
)

type ClientExample interface {
	WithContext(context.Context) ClientExample
	Context() context.Context
	CreateUserInfo(req *CreateUserInfo) (*GithubComShrewxToolxExampleRouterCrudCreateUserInfoResponse, error)
	DownloadFile(req *DownloadFile) (*GithubComShrewxGinxAttachment, error)
	GetHelloWorld(req *GetHelloWorld) (*GithubComShrewxToolxExampleRouterCrudGetHelloWorldResponse, error)
	GetUserInfo(req *GetUserInfo) (*GithubComShrewxToolxExampleRouterCrudGetUserInfoResponse, error)
	HTML(req *HTML) (*GithubComShrewxGinxHTML, error)
	Image(req *Image) (*GithubComShrewxGinxImagePNG, error)
	ModifyUserInfo(req *ModifyUserInfo) (*string, error)
	UploadFile(req *UploadFile) error
}

func NewClientExample(c github_com_shrewx_ginx.Client) *ClientExampleStruct {
	return &(ClientExampleStruct{
		Client: c,
	})
}

type ClientExampleStruct struct {
	Client github_com_shrewx_ginx.Client
	ctx    context.Context
}

func (c *ClientExampleStruct) WithContext(ctx context.Context) ClientExample {
	cc := new(ClientExampleStruct)
	cc.Client = c.Client
	cc.ctx = ctx
	return cc
}
func (c *ClientExampleStruct) Context() context.Context {
	if c.ctx != nil {
		return c.ctx
	}
	return context.Background()
}

func (c *ClientExampleStruct) CreateUserInfo(req *CreateUserInfo) (*GithubComShrewxToolxExampleRouterCrudCreateUserInfoResponse, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) DownloadFile(req *DownloadFile) (*GithubComShrewxGinxAttachment, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) GetHelloWorld(req *GetHelloWorld) (*GithubComShrewxToolxExampleRouterCrudGetHelloWorldResponse, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) GetUserInfo(req *GetUserInfo) (*GithubComShrewxToolxExampleRouterCrudGetUserInfoResponse, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) HTML(req *HTML) (*GithubComShrewxGinxHTML, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) Image(req *Image) (*GithubComShrewxGinxImagePNG, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) ModifyUserInfo(req *ModifyUserInfo) (*string, error) {
	return req.InvokeAndBind(c.Context(), c.Client)
}

func (c *ClientExampleStruct) UploadFile(req *UploadFile) error {
	return req.InvokeAndBind(c.Context(), c.Client)
}
