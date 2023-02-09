package json_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmendezm/drone-dispatch-controller/dto"
	"github.com/jmendezm/drone-dispatch-controller/internal/services"
)

func RegisterDrone(s services.DroneServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.Drone
		if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "bad parameters"})
			return
		}
		if err := s.RegisterDrone(ctx, &req); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatus(200)
	}
}

func LoadDrone(s services.DroneServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.LoadDrone
		if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "bad parameters"})
			return
		}
		if err := s.LoadDrone(ctx, &req); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatus(200)
	}
}

func GetDroneLoad(s services.DroneServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serialNumber := ctx.Param("serial_number")
		load, err := s.GetDroneLoad(ctx, serialNumber)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(200, load)
	}
}

func GetAvailableDrones(s services.DroneServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		drones, err := s.GetAvailableDrones(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(200, drones)
	}
}

func GetDroneBattery(s services.DroneServiceInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		serialNumber := ctx.Param("serial_number")
		battery, err := s.GetDroneBattery(ctx, serialNumber)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(200, gin.H{"battery_level": battery})
	}
}
