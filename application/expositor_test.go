package application_test

import (
	"fitness/application"
	"testing"
)

func TestAbastecerExpositor(t *testing.T) {
	//Arrange
	expositor := &application.Expositor{
		ID:          "1",
		Localizacao: "Ginasio A",
		Estoque:     make(map[int]int),
	}

	comando := application.AbastecerExpositor{
		ExpositorID: "1",
		ProdutoID:   1,
		Quantidade:  10,
	}

	//Act
	evento := comando.AbastecerExpositor(expositor)

	//Assert
	if evento.ExpositorID != "1" || evento.ProdutoID != 1 || evento.Quantidade != 10 {
		t.Error("O expositor não foi abastecido")
	}

	if expositor.Estoque[1] != 10 {
		t.Error("O estoque do expositor está incorrecto após o abastecimento")
	}
}

func TestAbastecerExpositor_DoisAbastecimentos(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "2",
		Localizacao: "Ginásio D",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "2",
		ProdutoID:   1,
		Quantidade:  10,
	}

	// Act
	comando.AbastecerExpositor(expositor)
	evento2 := comando.AbastecerExpositor(expositor)

	// Assert
	if evento2.Quantidade != 10 {
		t.Errorf("O segundo abasteciemto deveria adicionar 10 unidades")
	}

	if expositor.Estoque[1] != 20 {
		t.Errorf("O estoque do expositor deveria ser 20 após o segundo abastecimento")
	}
}

func TestAbastecerExpositor_ProdutoDiferente(t *testing.T) {
	//Arrange
	expositor := &application.Expositor{
		ID:          "12",
		Localizacao: "Ginasio C",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "12",
		ProdutoID:   12,
		Quantidade:  20,
	}

	//Act
	evento := comando.AbastecerExpositor(expositor)

	//Assert
	if evento.ProdutoID != 12 {
		t.Error("ID do produto errado")
	}

	if expositor.Estoque[12] != 20 {
		t.Error("O estoque do expositor deveria ser 20 para produto 12")
	}
}

func TestAbastecerExpositor_ExpositorVazio(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "45",
		Localizacao: "Ginásio H",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "45",
		ProdutoID:   14,
		Quantidade:  15,
	}

	// Act
	evento := comando.AbastecerExpositor(expositor)

	// Assert
	if evento.ExpositorID != "45" {
		t.Error("O evento deveria conter o ID do expositor 45")
	}

	if expositor.Estoque[14] != 15 {
		t.Errorf("O estoque do expositor deveria ser 15 para 0 produto %v", evento.ProdutoID)
	}
}
