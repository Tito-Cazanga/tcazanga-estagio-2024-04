package application

import (
	"fitness/domain"
)

type ConsumirProduto struct {
	ExpositorID int
	ProdutoID   int
	Quantidade  int
}

func (comando *ConsumirProduto) RegistrarConsumo(expositor *Expositor) *domain.ConsumoRegistrado {

	// Verificar se o expositor possui o produto em estoque
	quantidadeEstoque, funcionar := expositor.Estoque[comando.ProdutoID]
	if !funcionar {
		return nil
	}

	if quantidadeEstoque < comando.Quantidade {
		return nil
	}

	expositor.Estoque[comando.ProdutoID] -= comando.Quantidade

	return &domain.ConsumoRegistrado{
		GinásioID:  expositor.ID,
		ProdutoID:  comando.ProdutoID,
		Quantidade: comando.Quantidade,
	}
}
