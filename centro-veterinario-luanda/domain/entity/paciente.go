package domain

import (
	"errors"
	"time"
)

type Patient struct {
	ID         string
	Name       string
	AdmittedAt time.Time
}

func NewPatient(id, name string) (*Patient, error) {
	if id == "" || name == "" {
		return nil, errors.New("data de entrada do paciente invlaido")
	}
	
	return &Patient{
		ID:         id,
		Name:       name,
		AdmittedAt: time.Now(),
	}, nil
}
