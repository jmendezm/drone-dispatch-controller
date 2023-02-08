package main

import (
	"flag"
	"github.com/jmendezm/drone-dispatch-controller/server"
)

func main() {
	configFilePath := flag.String("config", "./config/config.json", "Configuration file path")
	flag.Parse()
	server.RunServer(*configFilePath)
}
