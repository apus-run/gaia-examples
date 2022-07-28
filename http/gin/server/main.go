package main

import (
	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/log"
	grpcserver "github.com/apus-run/gaia/transport/grpc"

	pb "github.com/apus-run/gaia/examples/http/gin/proto"
	"github.com/apus-run/gaia/examples/http/gin/server/service"
	"github.com/apus-run/gaia/middleware/recovery"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-server"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

func NewGRPCServer(svc *service.UserServiceServer) *grpcserver.Server {
	// grpc server
	grpcServer := grpcserver.NewServer(
		grpcserver.Address(":9000"),
		grpcserver.Middleware(
			recovery.Recovery(),
		),
	)

	pb.RegisterUserServiceServer(
		grpcServer,
		svc,
	)

	return grpcServer
}

func main() {
	userServiceServer := service.NewUserServiceServer()
	gs := NewGRPCServer(userServiceServer)

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(
			gs,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
