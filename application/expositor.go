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

func (comando *AbastecerExpositor) AbastecerExpositor(expositor *Expositor) *domain.ExpositorAbastecido {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}

	expositor.Estoque[comando.ProdutoID] += comando.Quantidade

	return &domain.ExpositorAbastecido{
		ExpositorID: comando.ExpositorID,
		ProdutoID:   comando.ProdutoID,
		Quantidade:  comando.Quantidade,
	}
}
