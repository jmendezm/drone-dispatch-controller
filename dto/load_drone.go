package dto

type LoadDrone struct {
	DroneSerialNumber string           `json:"drone_serial_number"`
	Items             []MedicationItem `json:"items"`
}
