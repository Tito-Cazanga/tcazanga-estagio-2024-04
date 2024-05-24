package domain

import (
	"errors"
)

// ConsumirProduto é o comando para consumir um produto de um expositor.
type ConsumirProduto struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

// ConsumoRealizado é o evento gerado após um produto ser consumido de um expositor.
type ConsumoRealizado struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

// NewConsumirProduto cria uma nova instância do comando ConsumirProduto, validando a quantidade.
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

// RegistrarConsumo executa o comando de consumir produto no expositor e retorna o evento gerado.
func (cmd *ConsumirProduto) RegistrarConsumo(expositor *Expositor) (*ConsumoRealizado, error) {
	quantidadeEstoque, existe := expositor.Estoque[cmd.ProdutoID]
	if !existe {
		return nil, errors.New("produto não encontrado no estoque")
	}

	if quantidadeEstoque < cmd.Quantidade {
		return nil, errors.New("estoque insuficiente para o produto")
	}

	// Deduzir a quantidade consumida do estoque
	expositor.Estoque[cmd.ProdutoID] -= cmd.Quantidade

	// Gerar evento de consumo
	evento := &ConsumoRealizado{
		ExpositorID: expositor.ID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}

	return evento, nil
}
