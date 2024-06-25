package application

import (

	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"
)
type PatientService struct {
	Repo domain.PatientRepository
}

func NewPatient(repo domain.PatientRepository) *PatientService{
	return &PatientService{}
}

func (v *PatientService) InternarPaciente(){

}