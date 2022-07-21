package main

import (
	"flag"
	"log"

	"github.com/apus-run/gaia/config"
)

var flagconf string

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
	flag.StringVar(&flagconf, "conf", "gaia.yaml", "config path, eg: -conf gaia.yaml")
}

func main() {
	flag.Parse()
	if err := config.Load(flagconf); err != nil {
		panic(err)
	}

	gc := config.Get("http.server.address")
	log.Printf("address: %s", gc)

	fgc := config.File("gaia").Get("grpc.server.address")
	log.Printf("address: %s", fgc)

	config.Watch()
}
