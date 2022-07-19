package main

import (
	"context"
	"fmt"

	"github.com/apus-run/gaia"
	pb "github.com/apus-run/gaia/examples/helloworld/api"
	"github.com/apus-run/gaia/log"
	"github.com/apus-run/gaia/transport/grpc"
)

var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("invalid argument %s", in.Name)
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func main() {
	s := &server{}

	grpcServer := grpc.NewServer(
		grpc.Address(":9000"),
	)

	pb.RegisterGreeterServer(
		grpcServer,
		s,
	)

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(grpcServer),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
