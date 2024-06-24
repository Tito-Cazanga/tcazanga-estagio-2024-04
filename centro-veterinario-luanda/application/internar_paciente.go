package application

import domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"

type VeterinariaServico struct{
		repo domain.VeterinariaRepositorio
}

func NovoInternamento(repo domain.VeterinariaRepositorio)*VeterinariaServico{
	return &VeterinariaServico{}

} 
func (v *VeterinariaServico)VerificarPacienteInternado(){

}
