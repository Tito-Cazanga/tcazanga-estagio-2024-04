package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
)

func TestInternarPaciente(t *testing.T) {
	t.Run("Deve internar paciente com dados valido", func(t *testing.T) {
		//arrange
		repo := inmem.NovoRepositorioemMemoriaPaciente()
		servico := application.NovoPaciente(repo)

		//act
		pacienteid := "0011"
		pacientenome := "Max"
		dados := servico.InternarPaciente(pacienteid, pacientenome)

		//assert
		if dados != nil {
			t.Errorf("Erro ao internar paciente: %v", dados)
		}

		paciente, dados := repo.EncontrarID(pacienteid)
		if dados != nil {
			t.Errorf("Erro ao encontrar paciente: %v", dados)
		}

		if paciente.Nome != pacientenome {
			t.Errorf("Nome do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", pacientenome, paciente.Nome)
		}
		if paciente.Status != "Internado" {
			t.Errorf("Status do paciente nao coincide com o esperado. Esperado: %s, Obtido: %s", "Internado", paciente.Status)
		}
	})
}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	/*t.Run("Verifica se salva os dados do paciente no repositorio", func(t *testing.T) {
		//arrange
		var paciente1 = testando.Paciente{ID: "001", Nome: "Zeca", DataDeEntrada: "2023/04/06"}
		paciente2 := testando.Paciente{ID: "002", Nome: "Zequinha", DataDeEntrada: "2023/06/04"}
		
		//act
		err1 := application.SalvaPaciente(paciente1)
		err2 := application.SalvaPaciente(paciente2)
		//assert
		if err1 != nil {
			t.Errorf("Erro %v %v", err1,err2)
			t.Fail()
		}
	})*/

