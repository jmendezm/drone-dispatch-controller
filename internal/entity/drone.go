package entity

type Drone struct {
	SerialNumber    string
	Model           DroneModel
	WeightLimit     float32
	BatteryCapacity float32
	State           DroneState
}
