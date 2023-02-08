package services

import (
	"context"
	"github.com/jmendezm/drone-dispatch-controller/dto"
)

type DroneServiceInterface interface {
	RegisterDrone(ctx context.Context, drone *dto.Drone) error
	LoadDrone(ctx context.Context, load *dto.LoadDrone) error
	GetDroneLoad(ctx context.Context, droneSerial string) ([]*dto.MedicationItem, error)
	GetAvailableDrones(ctx context.Context) ([]*dto.Drone, error)
	GetDroneBattery(ctx context.Context, droneSerial string) (float32, error)
}
