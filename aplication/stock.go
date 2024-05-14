package application

import (
	//"fmt"
	"time"
)

// Produto representa um item no estoque
type Produto struct {
	Nome        string
	Validade    time.Time
	Quantidade  int
}

// Verifica produtos próximos à data de validade e produtos expirados
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

// Remove produtos expirados do estoque
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