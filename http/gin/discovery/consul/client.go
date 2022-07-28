package consul

import (
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

// Config consul config
type Config struct {
	Address    string
	Scheme     string
	Datacenter string
	WaitTime   time.Duration
	Namespace  string
}

func New(cfg *Config) (*api.Client, error) {
	consulClient, err := api.NewClient(&api.Config{
		Address:    cfg.Address,
		Scheme:     cfg.Scheme,
		Datacenter: cfg.Datacenter,
		WaitTime:   cfg.WaitTime,
		Namespace:  cfg.Namespace,
	})

	if err != nil {
		log.Fatal(err)
	}

	return consulClient, nil
}
