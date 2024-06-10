package application

import "time"

type Produto struct { 
	Nome string
	Quantidade int 
	Validade time.Time
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
