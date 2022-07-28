package main

import (
	"flag"
	"log"

	"github.com/apus-run/gaia/pkg/config"
)

var flagConf string

// Defines the config JSON Field

// AppConfig app config
type AppConfig struct {
	Http struct {
		Server struct {
			Address string
			Name    string
		}
	}

	GRPC struct {
		Server struct {
			Address string
			Name    string
		}
	}
}

func init() {
	flag.StringVar(&flagConf, "conf", "gaia.yaml", "config path, eg: -conf gaia.yaml")
}

func main() {
	flag.Parse()
	if err := config.Load(flagConf); err != nil {
		panic(err)
	}

	gc := config.Get("http.server.address")
	log.Printf("address: %s", gc)

	fgc := config.File("gaia").Get("grpc.server.address")
	log.Printf("address: %s", fgc)

	// config.Watch()
}
