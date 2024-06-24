package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"
)

func Test_Internar_Paciente(t *testing.T){
	t.Run("Verificar se paciente foi internado", func(t *testing.T) {
		//arrange
		repo := domain.NovoInternamentoRepositorio()
		servico := application.NovoInternamento(repo)
		//act 
		servico.VerificarPacienteInternado()

	})
}