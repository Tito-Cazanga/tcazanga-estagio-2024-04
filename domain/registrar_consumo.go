package domain

import (
	"errors"
)

type ConsumirProduto struct {
	GinasioID string
	ProdutoID   int
	Quantidade  int
}

func NewConsumirProduto(ginasioID string, produtoID, quantidade int) (*ConsumirProduto, error) {
	if quantidade <= 0 {
		return nil, errors.New("a quantidade deve ser positiva")
	}

	return &ConsumirProduto{
		GinasioID:  ginasioID,
		ProdutoID:  produtoID,
		Quantidade: quantidade,
	}, nil
}



func (cmd *ConsumirProduto) RegistrarConsumo(expositor *Expositor) *ConsumirProduto {

	// Verificar se o expositor possui o produto em estoque
	quantidadeEstoque, funcionar := expositor.Estoque[cmd.ProdutoID]
	if !funcionar {
		return nil
	}

	if quantidadeEstoque < cmd.Quantidade {
		return nil
	}

	expositor.Estoque[cmd.ProdutoID] -= cmd.Quantidade

	return &ConsumirProduto{
		GinasioID:  expositor.ID,
		ProdutoID:  cmd.ProdutoID,
		Quantidade: cmd.Quantidade,
	}
}
