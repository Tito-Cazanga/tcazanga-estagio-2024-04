package stock

import (
	"reflect"
	"testing"
	"time"
)

func TestAdicionaProduto(t *testing.T) {
	// Arrange
	iv := NewInventario()

	// Act
	produto := &Produto{ID: 1, Descricao: "Leite", Categoria: "Laticionio", CondicaoArmazenamento: "Refrigerado"}
	iv.AdicionaProduto(produto)

	// Assert
	if len(iv.Produtos) != 1 || !reflect.DeepEqual(iv.Produtos[0], produto) {
		t.Errorf("Erro ao adicionar produto. Esperado %v, Recebido %v", produto, iv.Produtos)
	}

}


func TestAdicionaLote(t *testing.T) {
	// Arrange
	iv := NewInventario()

	// Act
	produto := &Produto{ID: 1, Descricao: "Lesite", Categoria: "Laticionio", CondicaoArmazenamento: "Refrigerado"}
	iv.AdicionaProduto(produto)

	lote := &Lote{Produto: produto, NumeroDoLote: "LOT123", DataDeEntrada: "12-04-12", DataDeExpiracao: time.Now(), Quantidade: 100}
	iv.AdicionaLote(lote)

	// Assert
	if len(iv.Lotes) != 1 || !reflect.DeepEqual(iv.Lotes[0], lote) {
		t.Errorf("Errro ao adicionar o lote. Esperado %v, Recebido %v", lote, iv.Lotes)
	}

}

