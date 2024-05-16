package application

import (
	"time"
)

type Produto struct { 
	Nome string
	Quantidade int 
	Validade time.Time
}

type Expositor struct {
	ID 				string
	Localizacao 	string
	Estoque 		map[int]int

}

type ExpositorAbastecido struct {
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}

type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}

type ConsumoRegistrado struct {
	GinásioID  string
	ProdutoID  int
	Quantidade int
}

type ConsumirProduto struct {
	ExpositorID int
	ProdutoID   int
	Quantidade  int
}

func (comando *AbastecerExpositor) AbastecerExpositor(expositor *Expositor) *ExpositorAbastecido {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}
	
	expositor.Estoque[comando.ProdutoID] += comando.Quantidade

	return &ExpositorAbastecido{
		ExpositorID: comando.ExpositorID,
		ProdutoID: comando.ProdutoID,
		Quantidade: comando.Quantidade,
	}
}

func VenderProduto(produtos *[]Produto, nomeProduto string, quantidadeVendida int) {
	for i, p := range *produtos {
		if p.Nome == nomeProduto {
			(*produtos)[i].Quantidade -= quantidadeVendida
			break
		}
	}
}

func VerificarProdutoNoExpositor(produtos *[]Produto, produtoEsperado Produto) bool {
	for _, p := range *produtos {
		if p.Nome == produtoEsperado.Nome && p.Quantidade == produtoEsperado.Quantidade {
			return true
		}
	}
	return false
}

func (comando *ConsumirProduto) ConsumirExpositor(expositor *Expositor) *ConsumoRegistrado {
	// Verificar se o expositor possui o produto em estoque
	quantidadeEstoque, ok := expositor.Estoque[comando.ProdutoID]
	if !ok {
		return nil
	}

	if quantidadeEstoque < comando.Quantidade {
		return nil
	}

	expositor.Estoque[comando.ProdutoID] -= comando.Quantidade

	return &ConsumoRegistrado{
		GinásioID:  expositor.ID,
		ProdutoID:  comando.ProdutoID,
		Quantidade: comando.Quantidade,
	}
}