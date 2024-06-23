package application_test

import (
	"testing"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/internal/application"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/internal/domain"
)

func Test_internamento_paciente(t *testing.T){
	t.Run("Verificar internamento de Paciente", func(t *testing.T){
		//arrange
		repo := domain.PatientHospitalization()
		service := application.NewVeterinaryService(repo)

		//act
		service.CheckHospitalization()

		//assert
		if repo == nil{
			t.Error("Não foi internando nenhum paciente")
		}
		
	})
}