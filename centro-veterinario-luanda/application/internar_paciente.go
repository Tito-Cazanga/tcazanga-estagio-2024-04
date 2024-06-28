package application

import (
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"

	entidade "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)

type PacienteServico struct {
	Repo domain.PacienteRepositorio
}

func NovoPaciente(repo domain.PacienteRepositorio) *PacienteServico {
	return &PacienteServico{Repo: repo}
}

func (v *PacienteServico) InternarPaciente(id, nomePaciente, raca string) error {
	paciente, err := entidade.NovoPaciente(id, nomePaciente, raca)
	if err != nil {
		return err
	}
	
	paciente.Status = "Internado"
	return v.Repo.Salvar(paciente)
}

func (v *PacienteServico) ConsultarPaciente(id string) (*entidade.Paciente, error) {
	repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes.csv")
	return repo.EncontrarID(id)
}
