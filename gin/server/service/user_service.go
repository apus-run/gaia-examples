package service

import (
	"context"

	pb "github.com/apus-run/gaia/examples/gin/proto"
)

var (
	_ pb.UserServiceServer = (*UserServiceServer)(nil)
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{}
}

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{
		Id:       1,
		Username: req.Username,
	}, nil
}
func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{
		Id:    1,
		Token: "xxxxxxxx",
	}, nil
}
func (s *UserServiceServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{
		Username: req.Username,
	}, nil
}
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserServiceServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}
func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{
		User: &pb.User{
			Id:       req.Id,
			Username: "moocss",
			Email:    "moocss@163.com",
		},
	}, nil
}
func (s *UserServiceServer) BatchGetUsers(ctx context.Context, req *pb.BatchGetUsersRequest) (*pb.BatchGetUsersReply, error) {
	return &pb.BatchGetUsersReply{}, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	if req.GetLimit() == 0 {
		req.Limit = 10
	}
	if req.GetPage() == 0 {
		req.Page = 1
	}
	return &pb.ListUserReply{
		Users: []*pb.User{
			{
				Id:       1,
				Username: "moocss",
				Email:    "moocss@163.com",
			},
			{
				Id:       2,
				Username: "moocss2",
				Email:    "moocss2@163.com",
			},
		},
		Page:  req.GetPage(),
		Limit: req.GetLimit() + 1,
		Total: 200,
	}, nil
}
