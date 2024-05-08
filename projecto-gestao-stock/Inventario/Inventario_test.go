package inventario_test

import (
	inventario "AcmeStock/Inventario"
	"testing"
	"time"
)

func TestNovoProdutoNaoNulo(t *testing.T) {
    dataDeValidade := time.Now().AddDate(0, 1, 0) 
    p := inventario.NovoProduto(1, "Produto Teste", dataDeValidade)

    if p == nil {
        t.Error("O produto não deve ser nulo.")
    }
}

func TestNovoProdutoID(t *testing.T) {
    dataDeValidade := time.Now().AddDate(0, 1, 0) 
    p := inventario.NovoProduto(1, "Produto Teste", dataDeValidade)

    if p.ID != 1 {
        t.Errorf("ID esperado: %d, ID retornado: %d", 1, p.ID)
    }
}
