// Code generated protoc-gen-go-gin. DO NOT EDIT.
// protoc-gen-go-gin v2.1.0

package proto

import (
	context "context"
	errcode "github.com/apus-run/gaia/pkg/errcode"
	ginx "github.com/apus-run/gaia/pkg/ginx"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the eagle package it is being compiled against.

// context.
// metadata.
// gin.ginx.errcode.

type UserServiceHTTPServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	ListUsers(context.Context, *ListUserRequest) (*ListUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterUserServiceHTTPServer(r gin.IRouter, srv UserServiceHTTPServer) {
	s := &UserService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type UserService struct {
	server UserServiceHTTPServer
	router gin.IRouter
}

func (s *UserService) Register_0_HTTP_Handler(ctx *ginx.Context) {
	var in RegisterRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).Register(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) Login_0_HTTP_Handler(ctx *ginx.Context) {
	var in LoginRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).Login(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) Logout_0_HTTP_Handler(ctx *ginx.Context) {
	var in LogoutRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).Logout(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) GetUser_0_HTTP_Handler(ctx *ginx.Context) {
	var in GetUserRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).GetUser(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) UpdateUser_0_HTTP_Handler(ctx *ginx.Context) {
	var in UpdateUserRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).UpdateUser(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) UpdatePassword_0_HTTP_Handler(ctx *ginx.Context) {
	var in UpdatePasswordRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).UpdatePassword(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) ListUsers_0_HTTP_Handler(ctx *ginx.Context) {
	var in ListUserRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		e := errcode.ErrInvalidParam.WithDetails(err.Error())
		ctx.Error(e)
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(UserServiceHTTPServer).ListUsers(newCtx, &in)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Success(out)
}

func (s *UserService) RegisterService() {
	s.router.Handle("POST", "/v1/auth/register", ginx.Handle(s.Register_0_HTTP_Handler))
	s.router.Handle("POST", "/v1/auth/login", ginx.Handle(s.Login_0_HTTP_Handler))
	s.router.Handle("POST", "/v1/auth/logout", ginx.Handle(s.Logout_0_HTTP_Handler))
	s.router.Handle("GET", "/v1/users/:id", ginx.Handle(s.GetUser_0_HTTP_Handler))
	s.router.Handle("PUT", "/v1/users/:id", ginx.Handle(s.UpdateUser_0_HTTP_Handler))
	s.router.Handle("PATCH", "/v1/users/password/:id", ginx.Handle(s.UpdatePassword_0_HTTP_Handler))
	s.router.Handle("GET", "/v1/users", ginx.Handle(s.ListUsers_0_HTTP_Handler))
}
