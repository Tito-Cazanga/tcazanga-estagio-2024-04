package inmem

import (
	"encoding/csv"
	"errors"
	"os"
	domain "github.com/Tito-Cazanga/tcazanga-estagio-2024-04/domain/entity"
)

type RepositorioemMemoriaPaciente struct {
	filePath string
}

func NovoRepositorioemMemoriaPaciente(filePath string) *RepositorioemMemoriaPaciente {
	return &RepositorioemMemoriaPaciente{
		filePath: filePath,
	}
}

func (r *RepositorioemMemoriaPaciente) Salvar(p *domain.Paciente) error {
	ficheiro, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer ficheiro.Close()

	escrever := csv.NewWriter(ficheiro)
	defer escrever.Flush()

	gravar := []string{p.ID, p.Nome, p.Raca, p.StatusInternamento}
	err = escrever.Write(gravar)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositorioemMemoriaPaciente) EncontrarID(id string) (*domain.Paciente, error) {
	ficheiro, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer ficheiro.Close()

	ler := csv.NewReader(ficheiro)
	gravars, err := ler.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, gravar := range gravars {
		if gravar[0] == id {
			return &domain.Paciente{
				ID:                 gravar[0],
				Nome:               gravar[1],
				Raca:               gravar[2],
				StatusInternamento: gravar[3],
			}, nil
		}
	}

	return nil, errors.New("paciente nao encontrado")
}
