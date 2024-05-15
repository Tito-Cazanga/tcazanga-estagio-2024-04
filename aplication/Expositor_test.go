package application_test

import (
	application "fitness/aplication"
	"testing"
)

func TestAbastecerExpositor(t *testing.T) {
	//Arrange
	expositor := &application.Expositor{ 
		ID:  "1",
		Localizacao: "Ginasio A",
		Estoque: make(map[int]int),
	}

	comando := application.AbastecerExpositor{
		ExpositorID: "1",
		ProdutoID: 1,
		Quantidade: 10,
	}

	//Act
	evento := comando.Executar(expositor)
	
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
		ID: "2",
		Localizacao: "Ginásio D",
		Estoque: make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "2",
		ProdutoID: 1,
		Quantidade: 10,
	}

	// Act
	comando.Executar(expositor)
	evento2 := comando.Executar(expositor)

	// Assert
	if evento2.Quantidade != 10 {
		t.Errorf("O segundo abasteciemto deveria adicionar 10 unidades")
	}

	if  expositor.Estoque[1] != 20 {
		t.Errorf("O estoque do expositor deveria ser 20 após o segundo abastecimento")
	}
}

func TestAbastecerExpositor_ProdutoDiferente(t *testing.T){
	//Arrange
	expositor := &application.Expositor{
		ID: "12",
		Localizacao: "Ginasio C",
		Estoque: make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "12",
		ProdutoID: 12,
		Quantidade: 20,
	}

	//Act
	evento := comando.Executar(expositor)
	
	//Assert
	if evento.ProdutoID != 12{
		t.Error("ID do produto errado")
	}

	if expositor.Estoque[12] != 20 {
		t.Error("O estoque do expositor deveria ser 20 para produto 12")
	}
}

