package domain

import (
	"errors"

	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)


type PatientRepository interface {
	Save(patient *domain.Patient) error
	FindByID(id string) (*domain.Patient, error)
}

type InMemoryPatientRepository struct {
	patients map[string]*domain.Patient
}

func NewInMemoryPatientRepository() *InMemoryPatientRepository {
	return &InMemoryPatientRepository{
		patients: make(map[string]*domain.Patient),
	}
}

func (r *InMemoryPatientRepository) Save(p *domain.Patient) error {
	if p.ID == "" {
		return errors.New("ID do paciente invalido")
	}
	r.patients[p.ID] = p
	return nil
}

func (r *InMemoryPatientRepository) FindByID(id string) (*domain.Patient, error) {
	p, exists := r.patients[id]
	if !exists {
		return nil, errors.New("paciente nao enontrado")
	}
	return p, nil
}

