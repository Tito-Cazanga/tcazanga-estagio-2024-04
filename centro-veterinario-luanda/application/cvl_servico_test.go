package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
)

func TestInternarPaciente(t *testing.T) {
	repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes_test.csv")
	servico := application.NovoPaciente(repo)

	

	t.Run("Deve internar paciente com dados válidos", func(t *testing.T) {
		internarPacienteComSucesso(t, *servico, "0011", "Max", "Rafeiro")
	})

	t.Run("Deve retornar erro ao internar paciente com ID inválido", func(t *testing.T) {
		internarPacienteComErro(t, *servico, "", "Max", "Rafeiro")
	})

	t.Run("Deve verificar se o paciente foi internado", func(t *testing.T) {
		verificarPacienteInternado(t, *repo, "0011", "Max", "Rafeiro")
	})
	
	t.Run("Internar Paciente já internado", func(t *testing.T) {
		err := servico.InternarPaciente("0011", "Max", "Rafeiro")
		if err == nil {
			t.Errorf("Erro esperado ao internar paciente com mesmo ID, mas nenhum erro foi retornado")
		}

		erroEsperado := "paciente com mesmo ID já existe"
		if err.Error() != erroEsperado {
			t.Errorf("Erro esperado: %s, obtido: %s", erroEsperado, err.Error())
		}
	})
}

func internarPacienteComSucesso(t *testing.T, servico application.PacienteServico,
	pacienteId string, pacienteNome string, pacienteRaca string) {
	err := servico.InternarPaciente(pacienteId, pacienteNome, pacienteRaca)

	if err != nil {
		t.Errorf("Erro ao internar paciente: %v", err)
	}
}

func internarPacienteComErro(t *testing.T, servico application.PacienteServico, pacienteId string, pacienteNome string, pacienteRaca string) {
	err := servico.InternarPaciente(pacienteId, pacienteNome, pacienteRaca)

	if err == nil {
		t.Errorf("Erro esperado ao internar paciente com dados inválidos, mas nenhum erro foi retornado")
	}
}

func verificarPacienteInternado(t *testing.T, repo inmem.RepositorioemMemoriaPaciente,
	pacienteId string, pacienteNome string, pacienteRaca string) {

	paciente, err := repo.EncontrarID(pacienteId)

	if err != nil {
		t.Errorf("Erro ao obter paciente: %v", err)
	}

	if paciente == nil {
		t.Errorf("Paciente não foi encontrado no repositório")
	}

	if paciente.ID != pacienteId {
		t.Errorf("Paciente encontrado não é o esperado. Esperado: %s, Obtido: %s", pacienteId, paciente.ID)
	}

	if paciente.Nome != pacienteNome {
		t.Errorf("Nome do paciente encontrado não é o esperado. Esperado: %s, Obtido: %s", pacienteNome, paciente.Nome)
	}

	if paciente.Raca != pacienteRaca {
		t.Errorf("Raça do paciente encontrado não é o esperado. Esperado: %s, Obtido: %s", pacienteRaca, paciente.Raca)
	}

}
