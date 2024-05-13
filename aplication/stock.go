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

		if diasParaExpirar < 7 && diasParaExpirar >= 0 {
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

/*
func main() {
	// Exemplo de lista de produtos (pode ser obtida de um banco de dados)
	produtos := []Produto{
		{Nome: "Whey Protein", Validade: time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC), Quantidade: 10},
		{Nome: "BCAA", Validade: time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC), Quantidade: 5},
		{Nome: "Creatina", Validade: time.Date(2024, time.August, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
        {Nome: "Creatina", Validade: time.Date(2024, time.May, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
	}

	// Verifica produtos próximos à data de validade e produtos expirados
	proximosValidade, expirados := verificarValidade(produtos)

	// Mostra produtos próximos à data de validade
	fmt.Println("Produtos próximos à data de validade:")
	for _, p := range proximosValidade {
		fmt.Printf("Produto: %s, Expira em: %s\n", p.Nome, p.Validade.Format("02/01/2006"))
	}

	// Mostra produtos expirados
	fmt.Println("\nProdutos expirados:")
	for _, p := range expirados {
		fmt.Printf("Produto: %s, Expirado em: %s\n", p.Nome, p.Validade.Format("02/01/2006"))
	}

	// Remove produtos expirados do estoque
	removerProdutosExpirados(&produtos)

	// Mostra produtos restantes após a remoção dos expirados
	fmt.Println("\nProdutos restantes no estoque após a remoção dos expirados:")
	for _, p := range produtos {
		fmt.Printf("Produto: %s, Expira em: %s\n", p.Nome, p.Validade.Format("02/01/2006"))
	}
}
*/