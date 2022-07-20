package main

import (
	"context"
	"fmt"
	"github.com/apus-run/gaia"
	pb "github.com/apus-run/gaia/examples/user-service/api"
	"github.com/apus-run/gaia/log"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-server"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	if in.Username == "" {
		return nil, fmt.Errorf("invalid argument %s", in.Username)
	}
	return &pb.CreateUserReply{
		Id:       10001,
		Username: in.Username,
		Email:    in.Email,
	}, nil
}

func main() {
	s := &server{}

	// grpc server
	grpcServer := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)

	pb.RegisterUserServiceServer(
		grpcServer,
		s,
	)

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(
			grpcServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
