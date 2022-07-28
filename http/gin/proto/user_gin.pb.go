// Code generated protoc-gen-go-gin. DO NOT EDIT.
// protoc-gen-go-gin v1.0.0

package proto

import (
	context "context"
	errcode "github.com/apus-run/gaia/pkg/errcode"
	xgin "github.com/apus-run/gaia/pkg/xgin"
	gin "github.com/gin-gonic/gin"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the eagle package it is being compiled against.

// context.
// metadata.
// gin.xgin.errcode.

type UserServiceHTTPServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterUserServiceHTTPServer(r gin.IRouter, srv UserServiceHTTPServer) {
	s := UserService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type UserService struct {
	server UserServiceHTTPServer
	router gin.IRouter
}

func (s *UserService) Register_0(ctx *xgin.Context) {
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

func (s *UserService) Login_0(ctx *xgin.Context) {
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

func (s *UserService) Logout_0(ctx *xgin.Context) {
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

func (s *UserService) GetUser_0(ctx *xgin.Context) {
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

func (s *UserService) UpdateUser_0(ctx *xgin.Context) {
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

func (s *UserService) UpdatePassword_0(ctx *xgin.Context) {
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

func (s *UserService) RegisterService() {
	s.router.Handle("POST", "/v1/auth/register", xgin.Handle(s.Register_0))
	s.router.Handle("POST", "/v1/auth/login", xgin.Handle(s.Login_0))
	s.router.Handle("POST", "/v1/auth/logout", xgin.Handle(s.Logout_0))
	s.router.Handle("GET", "/v1/users/:id", xgin.Handle(s.GetUser_0))
	s.router.Handle("PUT", "/v1/users/:id", xgin.Handle(s.UpdateUser_0))
	s.router.Handle("PATCH", "/v1/users/password/:id", xgin.Handle(s.UpdatePassword_0))
}
