package services

import (
	"context"
	"errors"
	"github.com/jmendezm/drone-dispatch-controller/dto"
	"github.com/jmendezm/drone-dispatch-controller/internal/entity"
	"github.com/jmendezm/drone-dispatch-controller/internal/repositories"
)

type DroneService struct {
	DBRepo repositories.DroneDBRepository
}

func (ds *DroneService) RegisterDrone(ctx context.Context, drone *dto.Drone) error {
	var err error
	var model entity.DroneModel
	var state entity.DroneState
	model, err = entity.DroneModelFromString(drone.Model)
	if err != nil {
		return err
	}
	state, err = entity.DroneStateFromString(drone.State)
	if err != nil {
		return err
	}
	d := &entity.Drone{
		SerialNumber:    drone.SerialNumber,
		Model:           model,
		WeightLimit:     drone.WeightLimit,
		BatteryCapacity: drone.BatteryCapacity,
		State:           state,
	}
	if err = d.Validate(); err != nil {
		return err
	}
	err = ds.DBRepo.RegisterDrone(ctx, d)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DroneService) LoadDrone(ctx context.Context, load *dto.LoadDrone) error {
	var err error
	items := make([]*entity.Medication, 0)
	for _, l := range load.Items {
		m := &entity.Medication{
			Name:   l.Name,
			Weight: l.Weight,
			Code:   l.Code,
			Image:  l.Image,
		}
		if err = m.Validate(); err != nil {
			return err
		}
		items = append(items, m)
	}
	var drone *entity.Drone
	drone, err = ds.DBRepo.GetDroneBySerialNumber(ctx, load.DroneSerialNumber)
	if err != nil {
		return err
	}
	if err = drone.LoadUp(items); err != nil {
		return err
	}
	err = ds.DBRepo.LoadDrone(ctx, drone)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DroneService) GetDroneLoad(ctx context.Context, droneSerial string) ([]*dto.MedicationItem, error) {
	var err error
	var items []*entity.Medication
	items, err = ds.DBRepo.GetDroneLoad(ctx, droneSerial)
	if err != nil {
		return nil, err
	}
	resp := make([]*dto.MedicationItem, 0)
	for _, i := range items {
		resp = append(resp, &dto.MedicationItem{
			Name:   i.Name,
			Weight: i.Weight,
			Code:   i.Code,
			Image:  i.Image,
		})
	}
	return resp, nil
}

func (ds *DroneService) GetAvailableDrones(ctx context.Context) ([]*dto.Drone, error) {
	var err error
	var avDrone []*entity.Drone
	avDrone, err = ds.DBRepo.GetAvailableDrones(ctx)
	if err != nil {
		return nil, err
	}
	resp := make([]*dto.Drone, 0)
	for _, i := range avDrone {
		resp = append(resp, &dto.Drone{
			SerialNumber:    i.SerialNumber,
			Model:           entity.DroneModelToString(i.Model),
			WeightLimit:     i.WeightLimit,
			BatteryCapacity: i.BatteryCapacity,
			State:           entity.DroneStateToString(i.State),
		})
	}
	return resp, nil
}

func (ds *DroneService) GetDroneBattery(ctx context.Context, droneSerial string) (float32, error) {
	var err error
	var drone *entity.Drone
	drone, err = ds.DBRepo.GetDroneBySerialNumber(ctx, droneSerial)
	if drone == nil {
		return 0, errors.New("drone not exists")
	}
	if err != nil {
		return 0, err
	}
	return drone.BatteryCapacity, nil
}
