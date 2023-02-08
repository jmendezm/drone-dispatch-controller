package entity

import "errors"

type DroneModel string

const (
	ModelLightweight   DroneModel = "Lightweight"
	ModelMiddleweight             = "Middleweight"
	ModelCruiserweight            = "Cruiserweight"
	ModelHeavyweight              = "Heavyweight"
)

func DroneModelToString(m DroneModel) string {
	return string(m)
}

func DroneModelFromString(m string) (DroneModel, error) {
	switch m {
	case string(ModelLightweight):
		return ModelLightweight, nil
	case ModelMiddleweight:
		return ModelMiddleweight, nil
	case ModelCruiserweight:
		return ModelCruiserweight, nil
	case ModelHeavyweight:
		return ModelHeavyweight, nil
	default:
		return "", errors.New("not valid drone model")
	}
}
