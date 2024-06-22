package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/internal/application"
)

func Test_internamento_paciente(t *testing.T){
	t.Run("Verificar internamento de Paciente", func(t *testing.T){
		//arrange
		service := application.NewVeterinaryService()


		//act
		service.CheckHospitalization()


		//assert
		
	})
}