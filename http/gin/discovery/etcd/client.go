package etcd

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// Config etcd config
type Config struct {
	Endpoints        []string
	BasicAuth        bool
	UserName         string
	Password         string
	ConnectTimeout   time.Duration // 连接超时时间
	Secure           bool
	AutoSyncInterval time.Duration // 自动同步member list的间隔
	TTL              int           // 单位：s
}

// Client ...
type Client struct {
	*clientv3.Client
	config *Config
}

// New ...
func New(config *Config) (*Client, error) {
	conf := clientv3.Config{
		Endpoints:            config.Endpoints,
		DialTimeout:          config.ConnectTimeout,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 3 * time.Second,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
			//grpc.WithUnaryInterceptor(grpcprom.UnaryClientInterceptor),
			//grpc.WithStreamInterceptor(grpcprom.StreamClientInterceptor),
		},
		AutoSyncInterval: config.AutoSyncInterval,
	}

	if config.Endpoints == nil {
		return nil, fmt.Errorf("[etcd]  client etcd endpoints empty, empty endpoints")
	}

	if !config.Secure {
		conf.DialOptions = append(conf.DialOptions, grpc.WithInsecure())
	}

	if config.BasicAuth {
		conf.Username = config.UserName
		conf.Password = config.Password
	}

	client, err := clientv3.New(conf)

	if err != nil {
		return nil, fmt.Errorf("[etcd] client etcd start failed: %v", err)
	}

	cc := &Client{
		Client: client,
		config: config,
	}

	return cc, nil
}
