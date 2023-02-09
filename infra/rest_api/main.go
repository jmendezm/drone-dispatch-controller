package http2server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmendezm/drone-dispatch-controller/config"
	"github.com/jmendezm/drone-dispatch-controller/infra/rest_api/handlers"
	"github.com/jmendezm/drone-dispatch-controller/internal/services"
	log "github.com/sirupsen/logrus"
	"io"
)

type RestAPI struct {
	Configuration *config.Config
	DroneService  services.DroneServiceInterface
	server        *gin.Engine
}

func (s *RestAPI) New() {
	s.server = gin.New()
	ginLogger := gin.Logger()
	if !s.Configuration.ShowLogs {
		ginLogger = gin.LoggerWithConfig(gin.LoggerConfig{
			Output: io.Discard,
		})
	}
	s.server.Use(ginLogger)
	s.server.Use(gin.Recovery())

	handlers.RegisterHandlers(s.server, s.DroneService)
}

func (s *RestAPI) Start() {
	if err := s.server.Run(fmt.Sprintf(":%d", s.Configuration.ListenPort)); err != nil {
		log.Fatal(err)
	}
}
