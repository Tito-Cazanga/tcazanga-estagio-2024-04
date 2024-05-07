package validade

import (
	"testing"

)

func TestCriaNovoProduto_ProdutoCriado(t *testing.T) {
	// Arrange
	iv := NewInventario()

	produto := &Produto{ID: "12ds", Descricao: "Sumol", Categoria: "Bebidas", AmbienteDeArmazenamento: "Refrigerado"}

	// Act
	iv.CriaNovoProduto(produto)

	// Assert
	if iv.Produtos[0].ID != produto.ID {
		t.Errorf("Esperado %v, Recebido %v: ", produto, iv.Produtos)
	}

}


func TestFalhaAoCriarNovoProduto(t *testing.T) {
		// Arrange
		iv := NewInventario()

		produto := &Produto{ID: "12ds", Descricao: "Sumol", Categoria: "Bebidas", AmbienteDeArmazenamento: "Refrigerado"}
	
		// Act
		iv.CriaNovoProduto(produto)
	
		// Assert
		if iv.Produtos[0].ID == produto.ID {
			t.Errorf("Esperado %v, Recebido %v: ", produto, iv.Produtos)
		}
	

}