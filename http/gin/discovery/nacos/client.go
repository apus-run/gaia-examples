package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// Config consul config
type Config struct {
	Address     string
	Port        uint64
	NamespaceId string // default public
	TimeoutMs   uint64 // unit: default 10000ms
	LogDir      string
	CacheDir    string
	LogLevel    string // default value is info
}

func New(cfg *Config) (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.Address, cfg.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         cfg.NamespaceId,
		TimeoutMs:           cfg.TimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              cfg.LogDir,
		CacheDir:            cfg.CacheDir,
		LogLevel:            cfg.LogLevel,
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
