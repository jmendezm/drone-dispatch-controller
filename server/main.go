package server

import (
	"github.com/jmendezm/drone-dispatch-controller/config"
	"github.com/jmendezm/drone-dispatch-controller/infra"
	restapi "github.com/jmendezm/drone-dispatch-controller/infra/rest_api"
	"github.com/jmendezm/drone-dispatch-controller/internal/repositories"
	"github.com/jmendezm/drone-dispatch-controller/internal/services"
)

func RunServer(configFilePath string) {
	conf := config.New(configFilePath)
	infra.InitLog(conf)

	DB := infra.SqliteDB{}
	DB.InitSqliteDB()
	DB.PrepareDB()

	Repo := &repositories.DroneSqliteRepository{
		Conn: DB.Conn,
	}

	service := &services.DroneService{
		DBRepo: Repo,
	}

	restAPI := restapi.RestAPI{
		Configuration: conf,
		DroneService:  service,
	}
	restAPI.New()
	restAPI.Start()
}
