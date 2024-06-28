package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
)

func TestInternarPaciente(t *testing.T) {
	t.Run("Deve internar paciente com dados valido", func(t *testing.T) {
		//arrange
		repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes.csv")
		servico := application.NovoPaciente(repo)

		//act
		pacienteId := "0011"
		pacienteNome := "Max"
		pacienteRaca := "Rafeiro"
		dados := servico.InternarPaciente(pacienteId, pacienteNome, pacienteRaca)

		//assert
		if dados!= nil {
			t.Errorf("Erro ao internar paciente: %v", dados)
		}

		paciente, dados := repo.EncontrarID(pacienteId)
		
		if dados!= nil {
			t.Errorf("Erro ao encontrar paciente: %v", dados)
		}

		if paciente.Nome != pacienteNome {
			t.Errorf("Nome do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", pacienteNome, paciente.Nome)
		}
		if paciente.Raca != pacienteRaca {
			t.Errorf("Raca do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", pacienteRaca, paciente.Raca)
		}
		if paciente.Status != "Internado" {
			t.Errorf("Status do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", "Internado", paciente.Status)
		}
	})
}


func TestInternarPaciente_com_id_invalido(t *testing.T) {
	// arrange
	repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes.csv")
	servico := application.NovoPaciente(repo)

	// act
	pacienteId := ""
	pacienteNome := "Max"
	pacienteRaca := "Rafeiro"
	err := servico.InternarPaciente(pacienteId, pacienteNome, pacienteRaca)

	// assert
	if err == nil {
		t.Errorf("Erro esperado ao internar paciente com dados inválidos, mas nenhum erro foi retornado")
	}
}