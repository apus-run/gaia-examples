package main

import (
	"context"
	"log"

	"github.com/apus-run/gaia"

	v1 "github.com/apus-run/gaia/examples/validate/api"
	"github.com/apus-run/gaia/middleware/validate"
	"github.com/apus-run/gaia/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "errors"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"
)

type server struct {
	v1.UnimplementedExampleServiceServer
}

func (s *server) TestValidate(ctx context.Context, in *v1.Request) (*v1.Reply, error) {
	return &v1.Reply{
		Message: "ok",
	}, nil
}

func main() {
	s := &server{}
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			validate.Validator(),
		))
	v1.RegisterExampleServiceServer(grpcSrv, s)

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithServer(
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
