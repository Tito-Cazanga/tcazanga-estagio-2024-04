package inmem

import (
	"errors"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)

type  RepositorioemMemoriaPaciente struct {
	pacientes map[string]*domain.Paciente
}

func NovoRepositorioemMemoriaPaciente() *RepositorioemMemoriaPaciente {
	return &RepositorioemMemoriaPaciente{
		pacientes: make(map[string]*domain.Paciente),
	}
}

func (r *RepositorioemMemoriaPaciente) Salvar(p *domain.Paciente) error {
	if p.ID == "" {
		return errors.New("ID do paciente invalido")
	}
	r.pacientes[p.ID] = p
	return nil
}

func (r *RepositorioemMemoriaPaciente) EncontrarID(id string) (*domain.Paciente, error) {
	p, exists := r.pacientes[id]
	if !exists {
		return nil, errors.New("paciente nao enontrado")
	}
	return p, nil
}

