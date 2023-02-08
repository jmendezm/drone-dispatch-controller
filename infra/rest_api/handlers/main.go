package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmendezm/drone-dispatch-controller/infra/rest_api/handlers/json_handlers"
	"github.com/jmendezm/drone-dispatch-controller/internal/services"
)

func RegisterHandlers(r *gin.Engine, srv services.DroneServiceInterface) {
	drone := r.Group("/drone")
	drone.POST("/", json_handlers.RegisterDrone(srv))
	drone.POST("/load", json_handlers.LoadDrone(srv))
	drone.GET("/load/:serial_number", json_handlers.GetDroneLoad(srv))
	drone.GET("/available", json_handlers.GetAvailableDrones(srv))
	drone.GET("/battery/:serial_number", json_handlers.GetDroneBattery(srv))
}
