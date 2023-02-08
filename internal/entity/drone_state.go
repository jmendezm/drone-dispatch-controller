package entity

import "errors"

type DroneState string

const (
	StateIDLE       DroneState = "IDLE"
	StateLOADING               = "LOADING"
	StateLOADED                = "LOADED"
	StateDELIVERING            = "DELIVERING"
	StateDELIVERED             = "DELIVERED"
	StateRETURNING             = "RETURNING"
)

func DroneStateToString(s DroneState) string {
	return string(s)
}

func DroneStateFromString(s string) (DroneState, error) {
	switch s {
	case string(StateIDLE):
		return StateIDLE, nil
	case StateLOADING:
		return StateLOADING, nil
	case StateLOADED:
		return StateLOADED, nil
	case StateDELIVERING:
		return StateDELIVERING, nil
	case StateDELIVERED:
		return StateDELIVERED, nil
	case StateRETURNING:
		return StateRETURNING, nil
	default:
		return "", errors.New("not valid drone state")
	}
}
