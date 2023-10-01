package main

import (
	"context"
	"net/http"
	"time"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/middleware"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/pkg/ginx"
	"github.com/apus-run/gaia/plugins/registry/consul"
	"github.com/apus-run/gaia/plugins/registry/nacos"
	"github.com/apus-run/gaia/registry"
	grpcserver "github.com/apus-run/gaia/transport/grpc"
	httpserver "github.com/apus-run/gaia/transport/http"
	"github.com/apus-run/sea-kit/log"
	"github.com/gin-gonic/gin"

	consulclient "github.com/apus-run/gaia/examples/gin/discovery/consul"
	nacosclient "github.com/apus-run/gaia/examples/gin/discovery/nacos"
	pb "github.com/apus-run/gaia/examples/gin/proto"
	"github.com/apus-run/gaia/examples/gin/web/service"

	_ "go.uber.org/automaxprocs"
)

var (
	// Name is the name of the compiled software.
	Name = "user-service-web"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

func ConnectGrpcServer() pb.UserServiceClient {
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

	c := pb.NewUserServiceClient(conn)

	return c
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
	g.Use(ginx.Middlewares(recovery.Recovery(), customMiddleware))

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

func getConsulDiscovery() registry.Discovery {
	client, err := consulclient.New(&consulclient.Config{
		Address:    "127.0.0.1:8500",
		Scheme:     "http",
		Datacenter: "",
		WaitTime:   5 * time.Millisecond,
		Namespace:  "",
	})
	if err != nil {
		panic(err)
	}
	return consul.New(client)
}

// TODO: nacos 的 endpoint 很特殊, 必须是 .grpc 后缀
// 例如: endpoint := "discovery:///user-service-server.grpc"
func getNacosDiscovery() registry.Discovery {
	client, err := nacosclient.New(&nacosclient.Config{
		Address:   "127.0.0.1",
		Port:      8848,
		TimeoutMs: 5000,
		LogDir:    "",
		CacheDir:  "warn",
	})

	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

func NewUserClient() pb.UserServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// consul 配置
	// endpoint := "discovery:///user-service-server"

	// nacos 配置
	endpoint := "discovery:///user-service-server.grpc"
	conn, err := grpcserver.DialInsecure(
		ctx,
		grpcserver.WithEndpoint(endpoint),
		grpcserver.WithDiscovery(getNacosDiscovery()),
	)
	if err != nil {
		panic(err)
	}
	c := pb.NewUserServiceClient(conn)
	return c
}

func main() {
	//c := ConnectGrpcServer()
	c := NewUserClient()
	userServiceServer := service.NewUserServiceServer(c)
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
