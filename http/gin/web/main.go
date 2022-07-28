package main

import (
	"context"
	"net/http"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/log"
	"github.com/apus-run/gaia/middleware"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/pkg/xgin"
	grpcserver "github.com/apus-run/gaia/transport/grpc"
	httpserver "github.com/apus-run/gaia/transport/http"
	"github.com/gin-gonic/gin"

	pb "github.com/apus-run/gaia/examples/http/gin/proto"
	"github.com/apus-run/gaia/examples/http/gin/web/service"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-web"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

var (
	UserServiceServerClient pb.UserServiceClient
)

func ConnectGrpcServer() {
	conn, err := grpcserver.DialInsecure(
		context.Background(),
		grpcserver.WithEndpoint("127.0.0.1:9000"),
		grpcserver.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	UserServiceServerClient = pb.NewUserServiceClient(conn)
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		log.Info("自定义标准插件")
		reply, err = handler(ctx, req)
		return
	}
}

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用gaia中间件
	g.Use(xgin.Middlewares(recovery.Recovery(), customMiddleware))

	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, map[string]string{"welcome": name})
	})

	return g
}

func NewHTTPServer(userSvc *service.UserServiceServer) *httpserver.Server {
	gin.SetMode("release")
	router := NewRouter()

	// http server
	httpServer := httpserver.NewServer(
		httpserver.Address(":8000"),
		httpserver.Middleware(
			recovery.Recovery(),
		),
	)

	httpServer.Handler = router

	pb.RegisterUserServiceHTTPServer(router, userSvc)

	return httpServer
}

func main() {
	ConnectGrpcServer()
	userServiceServer := service.NewUserServiceServer(UserServiceServerClient)
	hs := NewHTTPServer(userServiceServer)

	app := gaia.New(
		gaia.WithName(Name),
		gaia.WithVersion(Version),
		gaia.WithLogger(log.GetLogger()),
		gaia.WithServer(
			hs,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
