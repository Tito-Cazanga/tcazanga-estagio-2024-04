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
