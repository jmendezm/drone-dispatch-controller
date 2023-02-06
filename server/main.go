package server

import (
	"github.com/jmendezm/drone-dispatch-controller/config"
	"github.com/jmendezm/drone-dispatch-controller/infra"
	log "github.com/sirupsen/logrus"
)

func RunServer(configFilePath string) {
	conf := config.New(configFilePath)
	infra.InitLog(conf)
}
