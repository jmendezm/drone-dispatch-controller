package dto

type Drone struct {
	SerialNumber    string  `json:"serial_number"`
	Model           string  `json:"model"`
	WeightLimit     float32 `json:"weight_limit"`
	BatteryCapacity float32 `json:"battery_capacity"`
	State           string  `json:"state"`
}
