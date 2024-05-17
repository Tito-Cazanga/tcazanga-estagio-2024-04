package application_test

import (
	"fitness/application"
	"testing"
)

func TestConsumoRegistrado_AposExpositorAbastecido(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "1",
		Localizacao: "Ginásio A",
		Estoque:     map[int]int{1: 10},
	}
	comando := &application.ConsumirProduto{
		ExpositorID: 1,
		ProdutoID:   1,
		Quantidade:  5,
	}

	// Act
	evento := comando.RegistrarConsumo(expositor)

	// Assert
	if evento.GinásioID != "1" {
		t.Errorf("O ID do ginásio no evento está incorreto")
	}

	if evento.ProdutoID != 1 {
		t.Errorf("O ID do produto no evento está incorreto")
	}

	if evento.Quantidade != 5 {
		t.Errorf("A quantidade no evento está incorreta")
	}
}