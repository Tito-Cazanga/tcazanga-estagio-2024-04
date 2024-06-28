package domain

import (
	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)

type PacienteRepositorio interface {
	Salvar(paciente *domain.Paciente) error
	EncontrarID(id string) (*domain.Paciente, error)
}
