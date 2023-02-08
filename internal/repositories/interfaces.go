package repositories

import (
	"context"
	"github.com/jmendezm/drone-dispatch-controller/internal/entity"
)

type DroneDBRepository interface {
	GetDroneBySerialNumber(ctx context.Context, serialNumber string) (*entity.Drone, error)
	RegisterDrone(ctx context.Context, drone *entity.Drone) error
	LoadDrone(ctx context.Context, drone *entity.Drone) error
	GetAvailableDrones(ctx context.Context) ([]*entity.Drone, error)
	GetDroneLoad(ctx context.Context, droneSerial string) ([]*entity.Medication, error)
}
