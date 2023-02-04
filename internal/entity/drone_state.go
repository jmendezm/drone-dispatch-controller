package entity

type DroneState string

const (
	StatusIDLE       DroneState = "IDLE"
	StatusLOADING               = "LOADING"
	StatusLOADED                = "LOADED"
	StatusDELIVERING            = "DELIVERING"
	StatusDELIVERED             = "DELIVERED"
	StatusRETURNING             = "RETURNING"
)
