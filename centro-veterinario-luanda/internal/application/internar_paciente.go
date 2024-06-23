package application

import "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/internal/domain"

type VeterinaryService struct{

}

func NewVeterinaryService(repo *domain.NewVeterinaryRepository) *VeterinaryService{
	return &VeterinaryService{}
}

func (v *VeterinaryService) CheckHospitalization(){

}