package validade

import (
	"testing"
	"time"
)

func TestCriaNovoProduto_ProdutoCriado(t *testing.T) {
	// Arrange
	iv := NewInventario()

	produto := &Produto{ID: "12ds", Descricao: "Sumol", Categoria: "Bebidas", AmbienteDeArmazenamento: "Refrigerado"}

	// Act
	iv.CriaNovoProduto(produto)

	// Assert
	if iv.Produtos[0].ID != produto.ID {
		t.Errorf("Esperado %v, Recebido %v: ", produto, iv.Produtos[0])
	}

}

func TestFalhaAoCriarNovoProduto_ProdutoExistente(t *testing.T) {
	// Arrange
	iv := NewInventario()

	produto := &Produto{ID: "12ds", Descricao: "Leite", Categoria: "Laticinio", AmbienteDeArmazenamento: "Seco"}

	// Act
	iv.CriaNovoProduto(produto)

	// Assert
	if iv.Produtos[0].ID == produto.ID {
		t.Errorf("Esperado %v, Recebido %v: ", produto.ID, iv.Produtos[0].ID)
	}

}

func TestCriaNovoLote_LoteCriado(t *testing.T) {
	// Arrange
	iv := NewInventario()

	// Act
	lote := &Lote{ID: "12df", ProdutoId: "1", Lote: "LOT123", Validade: time.Now(), Quantidade: 100}
	iv.CriaNovoLote(lote)

	// Assert
	if iv.Lotes[0].ID != lote.ID  {
		t.Errorf("Esperado %v, Recebido %v", lote, iv.Lotes)
	}

}