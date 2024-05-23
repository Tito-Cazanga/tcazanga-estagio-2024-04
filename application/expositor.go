package application

import "fitness/domain"

type Expositor struct {
	ID          string
	Localizacao string
	Estoque     map[int]int
}
type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

func (cmd *AbastecerExpositor) AbastecerExpositor(expositor *Expositor) *domain.ExpositorAbastecido {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}

	expositor.Estoque[cmd.ProdutoID] += cmd.Quantidade

	return &domain.ExpositorAbastecido{
		ExpositorID: cmd.ExpositorID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}
}

