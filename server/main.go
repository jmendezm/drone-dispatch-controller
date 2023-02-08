package server

import (
	"github.com/jmendezm/drone-dispatch-controller/config"
	"github.com/jmendezm/drone-dispatch-controller/infra"
	restapi "github.com/jmendezm/drone-dispatch-controller/infra/rest_api"
)

func RunServer(configFilePath string) {
	conf := config.New(configFilePath)
	infra.InitLog(conf)

	restAPI := restapi.RestAPI{
		Configuration: conf,
		DroneService:  nil,
	}
	restAPI.New()
	restAPI.Start()
}
