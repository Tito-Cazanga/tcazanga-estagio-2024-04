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
		pacienteid := "0011"
		pacientenome := "Max"
		pacienteraca := "Rafeiro"
		dados := servico.InternarPaciente(pacienteid, pacientenome, pacienteraca)

		//assert
		if dados!= nil {
			t.Errorf("Erro ao internar paciente: %v", dados)
		}

		paciente, dados := repo.EncontrarID(pacienteid)
		if dados!= nil {
			t.Errorf("Erro ao encontrar paciente: %v", dados)
		}

		if paciente.Nome!= pacientenome {
			t.Errorf("Nome do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", pacientenome, paciente.Nome)
		}
		if paciente.Raca!= pacienteraca {
			t.Errorf("Raca do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", pacienteraca, paciente.Raca)
		}
		if paciente.Status!= "Internado" {
			t.Errorf("Status do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", "Internado", paciente.Status)
		}
	})
}
