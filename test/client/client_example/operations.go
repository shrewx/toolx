package client_example

import (
	context "context"

	github_com_shrewx_ginx "github.com/shrewx/ginx"
)

type CreateUserInfo struct {
	Authorization string `in:"header" name:"Authorization"`
	// Body
	Data struct {
		// 地址
		Address string
		// 年龄
		Age int32
		// 城市
		City GithubComShrewxToolxExampleConstantTypesCity
		// 职业
		Job GithubComShrewxToolxExampleConstantTypesJob
		// 名称
		Name string
	} `in:"body"`
}

func (CreateUserInfo) Path() string {
	return "/v0/crud"
}

func (CreateUserInfo) Method() string {
	return "POST"
}

func (req *CreateUserInfo) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *CreateUserInfo) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxToolxExampleRouterCrudCreateUserInfoResponse, error) {
	resp := new(GithubComShrewxToolxExampleRouterCrudCreateUserInfoResponse)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type DownloadFile struct {
	Authorization string `in:"header" name:"Authorization"`
}

func (DownloadFile) Path() string {
	return "/v0/file"
}

func (DownloadFile) Method() string {
	return "GET"
}

func (req *DownloadFile) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *DownloadFile) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxGinxAttachment, error) {
	resp := new(GithubComShrewxGinxAttachment)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type GetHelloWorld struct {
	Authorization string `in:"header" name:"Authorization"`
}

func (GetHelloWorld) Path() string {
	return "/v0/crud/hello"
}

func (GetHelloWorld) Method() string {
	return "GET"
}

func (req *GetHelloWorld) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *GetHelloWorld) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxToolxExampleRouterCrudGetHelloWorldResponse, error) {
	resp := new(GithubComShrewxToolxExampleRouterCrudGetHelloWorldResponse)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type GetUserInfo struct {
	Authorization string `in:"header" name:"Authorization"`
	Username      string `in:"query" name:"username"`
	ID            int32  `in:"path" name:"id"`
}

func (GetUserInfo) Path() string {
	return "/v0/crud/:id"
}

func (GetUserInfo) Method() string {
	return "GET"
}

func (req *GetUserInfo) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *GetUserInfo) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxToolxExampleRouterCrudGetUserInfoResponse, error) {
	resp := new(GithubComShrewxToolxExampleRouterCrudGetUserInfoResponse)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type HTML struct {
	Authorization string `in:"header" name:"Authorization"`
}

func (HTML) Path() string {
	return "/v0/file/html"
}

func (HTML) Method() string {
	return "GET"
}

func (req *HTML) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *HTML) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxGinxHTML, error) {
	resp := new(GithubComShrewxGinxHTML)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type Image struct {
	Authorization string `in:"header" name:"Authorization"`
}

func (Image) Path() string {
	return "/v0/file/image"
}

func (Image) Method() string {
	return "GET"
}

func (req *Image) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *Image) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*GithubComShrewxGinxImagePNG, error) {
	resp := new(GithubComShrewxGinxImagePNG)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type ModifyUserInfo struct {
	Authorization string `in:"header" name:"Authorization"`
	Name          string `in:"urlencoded" name:"name"`
}

func (ModifyUserInfo) Path() string {
	return "/v0/crud/:id"
}

func (ModifyUserInfo) Method() string {
	return "PUT"
}

func (req *ModifyUserInfo) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *ModifyUserInfo) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) (*string, error) {
	resp := new(string)

	response, err := req.Invoke(ctx, c)
	response.Bind(resp)

	return resp, err
}

type UploadFile struct {
	Authorization string                                `in:"header" name:"Authorization"`
	File1         *github_com_shrewx_ginx.MultipartFile `in:"multipart" name:"file1"`
}

func (UploadFile) Path() string {
	return "/v0/file"
}

func (UploadFile) Method() string {
	return "POST"
}

func (req *UploadFile) Invoke(ctx context.Context, c github_com_shrewx_ginx.Client) (github_com_shrewx_ginx.Response, error) {
	return c.Invoke(ctx, req)
}

func (req *UploadFile) InvokeAndBind(ctx context.Context, c github_com_shrewx_ginx.Client) error {

	_, err := req.Invoke(ctx, c)

	return err
}
