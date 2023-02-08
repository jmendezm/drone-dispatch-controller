package entity

import (
	"errors"
	"regexp"
)

type Medication struct {
	Name   string
	Weight float32
	Code   string
	Image  string
}

func (m *Medication) Validate() error {
	validName, _ := regexp.Match(`(\w|\d|_|-)+`, []byte(m.Name))
	validCode, _ := regexp.Match(`([A-Z]|\d|_|)+`, []byte(m.Code))
	if !validName {
		return errors.New("not valid medication name")
	}
	if !validCode {
		return errors.New("not valid medication code")
	}
	if m.Weight <= 0 {
		return errors.New("not valid medication weight")
	}
	return nil
}
