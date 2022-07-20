package main

import (
	"context"
	"fmt"
	hp "net/http"

	"github.com/apus-run/gaia"
	pb "github.com/apus-run/gaia/examples/helloworld/api"
	"github.com/apus-run/gaia/log"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
	"github.com/apus-run/gaia/transport/http"
	"github.com/gin-gonic/gin"
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

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(gin.Recovery())

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(hp.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return g
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

	pb.RegisterGreeterServer(
		grpcServer,
		s,
	)

	// http server
	httpServer := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			recovery.Recovery(),
		),
	)

	router := NewRouter()
	httpServer.Handler = router

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(
			grpcServer,
			httpServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
