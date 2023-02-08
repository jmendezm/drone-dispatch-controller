package entity

import "errors"

type Drone struct {
	SerialNumber    string
	Model           DroneModel
	WeightLimit     float32
	BatteryCapacity float32
	State           DroneState
	Load            []*Medication
}

func (d *Drone) Validate() error {
	if len(d.SerialNumber) > 100 {
		return errors.New("too long serial number")
	}
	if d.WeightLimit > 500 {
		return errors.New("too big weight limit")
	}
	if d.BatteryCapacity < 0 || d.BatteryCapacity > 100 {
		return errors.New("no valid battery capacity")
	}
	if d.SerialNumber == "" {
		return errors.New("no valid serial number")
	}
	if d.WeightLimit <= 0 {
		return errors.New("no valid weight limit")
	}
	return nil
}

func (d *Drone) LoadUp(items []*Medication) error {
	totalWeight := float32(0)
	for _, i := range items {
		totalWeight += i.Weight
	}

	if totalWeight > d.WeightLimit {
		return errors.New("too much weight to carry for the drone")
	}

	if d.BatteryCapacity < 25 {
		return errors.New("drone battery level is too low")
	}

	if d.State != StateIDLE {
		return errors.New("drone is not available right now")
	}

	d.State = StateLOADING
	d.Load = items
	return nil
}
