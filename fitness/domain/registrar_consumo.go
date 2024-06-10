package domain

import (
	"errors"
)

type ConsumirProduto struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

type ConsumoRealizado struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

func NewConsumirProduto(expositorID string, produtoID, quantidade int) (*ConsumirProduto, error) {
	if quantidade <= 0 {
		return nil, errors.New("a quantidade deve ser positiva")
	}

	return &ConsumirProduto{
		ExpositorID: expositorID,
		ProdutoID:   produtoID,
		Quantidade:  quantidade,
	}, nil
}

func (cmd *ConsumirProduto) RegistrarConsumo(expositor *Expositor) (*ConsumoRealizado, error) {
	quantidadeEstoque, existe := expositor.Estoque[cmd.ProdutoID]
	if !existe {
		return nil, errors.New("produto não encontrado no estoque")
	}

	if quantidadeEstoque < cmd.Quantidade {
		return nil, errors.New("estoque insuficiente para o produto")
	}

	expositor.Estoque[cmd.ProdutoID] -= cmd.Quantidade

	evento := &ConsumoRealizado{
		ExpositorID: expositor.ID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}

	return evento, nil
}
