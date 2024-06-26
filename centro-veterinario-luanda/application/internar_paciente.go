package application

import (
	"fmt"
	"os"

	"path/filepath"

	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/repositorio"

	entidade "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)

type PacienteServico struct {
	Repo domain.PacienteRepositorio
}

func NovoPaciente(repo domain.PacienteRepositorio) *PacienteServico {
	return &PacienteServico{}
}

func (v *PacienteServico) InternarPaciente(paciente string) {

}

// Função que verifica se paciente com o ID existe
func PacienteExiste(id string) bool {
	filePath := filepath.Join("/home/trainee/Desktop/tcazanga-estagio-2024-04/centro-veterinario-luanda/domain/repositorio",fmt.Sprintf("Paciente%s.txt", id))
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// Função que salva os dados do paciente num ficheiro de texto
func SalvaPaciente(paciente entidade.Paciente) error {
	// Verifica se o Paciente já existe
	if PacienteExiste(paciente.ID) {
		return fmt.Errorf("paciente com o ID %s já existe", paciente.ID)
	}

	// Cria ou abre um ficheiro
	filePath := filepath.Join("/home/trainee/Desktop/tcazanga-estagio-2024-04/centro-veterinario-luanda/domain/repositorio",fmt.Sprintf("Paciente%s.txt", paciente.Nome))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("falhou em criar o ficheiro: %v", err)
	}
	defer file.Close()

	// Escreve os dados do paciente no ficheiro
	_, err = fmt.Fprintf(file, "ID: %s\nNome: %s\nDataDeEntrada %s\n", paciente.ID, paciente.Nome, paciente.DataDeEntrada)
	if err != nil {
		return fmt.Errorf("falhou em escrever no ficheiro: %v", err)
	}

	fmt.Println("Dados do paciente salvos com sucesso.")
	return nil
}
