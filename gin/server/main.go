package main

import (
	"time"

	"github.com/apus-run/gaia"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/plugins/registry/consul"
	"github.com/apus-run/gaia/plugins/registry/nacos"
	"github.com/apus-run/gaia/registry"
	grpcserver "github.com/apus-run/gaia/transport/grpc"
	"github.com/apus-run/sea-kit/log"

	consulclient "github.com/apus-run/gaia/examples/gin/discovery/consul"
	nacosclient "github.com/apus-run/gaia/examples/gin/discovery/nacos"
	pb "github.com/apus-run/gaia/examples/gin/proto"
	"github.com/apus-run/gaia/examples/gin/server/service"
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

func getConsulRegistry() registry.Registry {
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

func getNacosRegistry() registry.Registry {
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
		// gaia.WithRegistry(getConsulRegistry()),
		gaia.WithRegistry(getNacosRegistry()),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
