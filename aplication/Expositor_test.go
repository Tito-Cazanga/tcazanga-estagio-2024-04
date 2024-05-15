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