package dto

type MedicationItem struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
	Code   string  `json:"code"`
	Image  string  `json:"image"`
}
