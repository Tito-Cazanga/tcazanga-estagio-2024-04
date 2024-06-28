package domain

import (
	"errors"
	"time"
)

const PacienteInternado = "Internado"
const RondaRealizado = "Ronda realizada"

type Paciente struct {
	ID                 string
	Nome               string
	Raca               string
	DataDeEntrada      time.Time
	StatusInternamento string
	StatusRonda			string
}




func NovoPaciente(id, nome, raca string) (*Paciente, error) {
	if id == "" || nome == "" || raca == "" {
		return nil, errors.New("data de entrada do paciente invalido")
	}

	return &Paciente{
		ID:                 id,
		Nome:               nome,
		Raca:               raca,
		DataDeEntrada:      time.Now(),
		StatusInternamento: PacienteInternado,
	}, nil
}

