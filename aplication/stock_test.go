package application

import (
	"testing"
	"time"
)

func TestVerificarValidade_ProximosEExpirados(t *testing.T) {
	// Arrange
	produtos := []Produto{
		{Nome: "Whey Protein", Validade: time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC), Quantidade: 10},
		{Nome: "BCAA", Validade: time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC), Quantidade: 5},
		{Nome: "Creatina", Validade: time.Date(2024, time.August, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
		{Nome: "Creatina", Validade: time.Date(2024, time.May, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
	}

	// Act
	proximosValidade, expirados := verificarValidade(produtos)

	// Assert
	if len(proximosValidade) != 2 {
		t.Errorf("Esperado 2 produtos próximos à validade, mas encontrados %d", len(proximosValidade))
	}

	// Assert
	if len(expirados) != 1 {
		t.Errorf("Esperado 1 produto expirado, mas encontrados %d", len(expirados))
	}
}

func TestRemoverExpirados(t *testing.T) {
	// Arrange
	produtos := []Produto{
		{Nome: "BCAA", Validade: time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC), Quantidade: 5},
		{Nome: "Creatina", Validade: time.Date(2024, time.August, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
	}

	// Act
	removerProdutosExpirados(&produtos)

	// Assert
	if len(produtos) != 2 {
		t.Errorf("Esperado 2 produtos, mas encontrados %d", len(produtos))
	}

	// Assert
	esperados := []Produto{
		{Nome: "BCAA", Validade: time.Date(2024, time.May, 20, 0, 0, 0, 0, time.UTC), Quantidade: 5},
		{Nome: "Creatina", Validade: time.Date(2024, time.August, 10, 0, 0, 0, 0, time.UTC), Quantidade: 8},
	}
	for i, p := range produtos {
		if p != esperados[i] {
			t.Errorf("Produto %d incorreto. Esperado: %v, Encontrado: %v", i, esperados[i], p)
		}
	}
}

func TestVerificarValidade_ProdutosComDatasDistantes(t *testing.T) {
	// Arrange: Definindo uma lista de produtos com datas de validade distantes
	produtos := []Produto{
		{Nome: "Produto B", Validade: time.Now().Add(60 * 24 * time.Hour), Quantidade: 5},   // Dentro de 60 dias
		{Nome: "Produto C", Validade: time.Now().Add(90 * 24 * time.Hour), Quantidade: 8},   // Dentro de 90 dias
		{Nome: "Produto D", Validade: time.Now().Add(120 * 24 * time.Hour), Quantidade: 12}, // Dentro de 120 dias
	}

	// Act: Executando a função verificarValidade
	proximosValidade, expirados := verificarValidade(produtos)

	// Assert: Verificando se não há produtos próximos à validade ou expirados
	if len(proximosValidade) != 0 {
		t.Errorf("Esperado nenhum produto próximo à validade, mas encontrados %d", len(proximosValidade))
	}

	if len(expirados) != 0 {
		t.Errorf("Esperado nenhum produto expirado, mas encontrados %d", len(expirados))
	}
}

func TestVerificarValidade_ProdutosComMesmaDataValidade(t *testing.T) {
	// Arrange: Definindo uma lista de produtos com a mesma data de validade
	dataValidade := time.Now().Add(10 * 24 * time.Hour) // Dentro de 10 dias
	produtos := []Produto{
		{Nome: "Produto A", Validade: dataValidade, Quantidade: 5},
		{Nome: "Produto B", Validade: dataValidade, Quantidade: 8},
		{Nome: "Produto C", Validade: dataValidade, Quantidade: 12},
	}

	// Act: Executando a função verificarValidade
	proximosValidade, _ := verificarValidade(produtos)

	// Assert: Verificando se todos os produtos estão corretamente identificados como próximos à validade
	if len(proximosValidade) != 3 {
		t.Errorf("Esperado 3 produtos próximos à validade, mas encontrados %d", len(proximosValidade))
	}
}

func TestRemoverExpirados_TodosProdutosExpirados(t *testing.T) {
	// Arrange: Definindo uma lista de produtos todos expirados
	produtos := []Produto{
		{Nome: "Produto A", Validade: time.Now().Add(-24 * time.Hour), Quantidade: 0},
		{Nome: "Produto B", Validade: time.Now().Add(-48 * time.Hour), Quantidade: 0},
		{Nome: "Produto C", Validade: time.Now().Add(-72 * time.Hour), Quantidade: 0},
	}

	// Act: Executando a função removerExpirados
	removerProdutosExpirados(&produtos)

	// Assert: Verificando se todos os produtos foram removidos
	if len(produtos) != 0 {
		t.Errorf("Esperado nenhum produto após remoção, mas encontrados %d", len(produtos))
	}
}

func TestVerificarValidade_AdicionarNovosProdutos(t *testing.T) {
	// Arrange: Definindo uma lista inicial de produtos
	produtos := []Produto{
		{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 10},
		{Nome: "Produto B", Validade: time.Now().Add(10 * 24 * time.Hour), Quantidade: 5},
	}

	// Act: Adicionando novos produtos à lista
	produtos = append(produtos, Produto{Nome: "Novo Produto", Validade: time.Now().Add(15 * 24 * time.Hour), Quantidade: 3})

	// Act: Executando a função verificarValidade após a adição de novos produtos
	proximosValidade, _ := verificarValidade(produtos)

	// Assert: Verificando se a quantidade de produtos próximos à validade está correta após a adição
	if len(proximosValidade) != 3 {
		t.Errorf("Esperado 3 produtos próximos à validade, mas encontrados %d", len(proximosValidade))
	}
}

func TestVerificarValidade_ListaVazia(t *testing.T) {
	// Arrange: Definindo uma lista de produtos vazia
	var produtos []Produto

	// Act: Executando a função verificarValidade
	proximosValidade, expirados := verificarValidade(produtos)

	// Assert: Verificando se não há produtos próximos à validade ou expirados
	if len(proximosValidade) != 0 {
		t.Errorf("Esperado nenhum produto próximo à validade, mas encontrados %d", len(proximosValidade))
	}

	if len(expirados) != 0 {
		t.Errorf("Esperado nenhum produto expirado, mas encontrados %d", len(expirados))
	}
}
