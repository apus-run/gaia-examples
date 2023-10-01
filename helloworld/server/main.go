package main

import (
	"context"
	"fmt"
	hp "net/http"
	"os"

	"github.com/apus-run/gaia"
	pb "github.com/apus-run/gaia/examples/helloworld/api"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
	"github.com/apus-run/gaia/transport/http"
	"github.com/apus-run/sea-kit/log"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	Version = "v1.0.0"

	id, _ = os.Hostname()
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

func Hello(c *gin.Context) {
	c.JSON(hp.StatusOK, gin.H{
		"message": "pong",
	})
}

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(gin.Recovery())

	g.GET("/", func(c *gin.Context) {
		c.String(hp.StatusOK, "ok")
	})

	g.GET("/hello", Hello)

	return g
}

func NewEchoRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(hp.StatusOK, "Hello, World!")
	})

	e.GET("/hello", func(c echo.Context) error {
		u := &User{
			Name:  "Kami",
			Email: "Kami@moocss.com",
		}
		return c.JSON(hp.StatusOK, u)
	})

	return e
}

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	)

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

	//router := NewEchoRouter()
	//httpServer.Handler = router

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(logger),
		gaia.WithServer(
			grpcServer,
			httpServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
