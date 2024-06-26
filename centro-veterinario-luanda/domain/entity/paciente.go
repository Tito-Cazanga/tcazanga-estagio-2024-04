package domain

import (
	/*"errors"
	"time"*/
)


const PacienteInternado = "Internado"

type Paciente struct {
	ID         		string
	Nome       		string
	DataDeEntrada 	string
}

/*func NovoPaciente(id, nome string) (*Paciente, error) {
	if id == "" || nome == "" {
		return nil, errors.New("data de entrada do paciente invalido")
	}
	
	return &Paciente{
		ID:         id,
		Nome:       nome,
		DataDeEntrada: time.Now(),
	}, nil
}*/
