package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"
)

func TestInternarPaciente(t *testing.T){
	t.Run("Deve internar paciente com dados validso", func(t *testing.T) {
		//arrange
		repo := domain.NewInMemoryPatientRepository()
		service := application.NewPatient(repo)
		
		//act
		service.InternarPaciente()


		//assert
		
	})
}