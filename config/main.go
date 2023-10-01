package main

import (
	"flag"
	"log"

	"github.com/apus-run/sea-kit/config"
	"github.com/apus-run/sea-kit/config/file"
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

	c := config.New(
		config.WithSource(
			// 添加前缀为 WEBSERVER_ 的环境变量，不需要的话也可以设为空字符串
			// env.NewSource("WEBSERVER_"),
			file.NewSource(flagConf),
		),
	)

	defer c.Close()

	// 加载配置源：
	if err := c.Load(); err != nil {
		log.Fatal(err)
	}

	var cf AppConfig
	if err := c.Scan(&cf); err != nil {
		panic(err)
	}
	log.Printf("配置文件: %+v", cf)

}
