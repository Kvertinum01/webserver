package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Kvertinum01/webserver/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/webserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := &server.Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.StartServer(config); err != nil {
		log.Fatal(err)
	}
}
