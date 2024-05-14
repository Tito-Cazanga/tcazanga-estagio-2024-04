package application

import (
	//"fmt"
	"time"
)
type Produto struct {
	Nome        string
	Validade    time.Time
	Quantidade  int
}


func verificarValidade(produtos []Produto) (proximosValidade []Produto, expirados []Produto) {
	tempoAtual := time.Now()

	for _, p := range produtos {
		diasParaExpirar := int(p.Validade.Sub(tempoAtual).Hours() / 24)

		if diasParaExpirar < 31 && diasParaExpirar >= 0 {
			proximosValidade = append(proximosValidade, Produto{Nome: p.Nome, Validade: p.Validade, Quantidade: p.Quantidade})
		} else if diasParaExpirar < 0 {
			expirados = append(expirados, Produto{Nome: p.Nome, Validade: p.Validade, Quantidade: p.Quantidade})
		}
	}

	return proximosValidade, expirados
}

func removerProdutosExpirados(produtos *[]Produto) {
	tempoAtual := time.Now()

	// Filtrando produtos não expirados
	produtosFiltrados := (*produtos)[:0]
	for _, p := range *produtos {
		if p.Validade.After(tempoAtual) {
			produtosFiltrados = append(produtosFiltrados, p)
		}
	}
	*produtos = produtosFiltrados
}
// ExisteProduto verifica se um produto existe na lista de produtos
func ExisteProduto(produtos []Produto, nomeProduto string) bool {
	for _, p := range produtos {
		if p.Nome == nomeProduto {
			return true
		}
	}
	return false
}
