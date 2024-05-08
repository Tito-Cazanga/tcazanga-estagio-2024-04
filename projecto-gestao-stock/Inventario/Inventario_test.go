package inventario_test

import (
	inventario "AcmeStock/Inventario"
	"testing"
	"time"
)

func TestNovoProduto(t *testing.T) {
	dataDeValidade := time.Now().AddDate(0, 1, 0) // Adiciona um mês à data atual
	p := inventario.NovoProduto(1, "Produto Teste", dataDeValidade)

	if p == nil {
		t.Error("O produto não deve ser nulo.")
	}

	if p.ID != 1 {
		t.Errorf("ID esperado: %d, ID retornado: %d", 1, p.ID)
	}

	if p.Nome != "Produto Teste" {
		t.Errorf("Nome esperado: %s, Nome retornado: %s", "Produto Teste", p.Nome)
	}

	if p.DataDeValidade != dataDeValidade {
		t.Errorf("Data de validade esperada: %s, Data de validade retornada: %s", dataDeValidade, p.DataDeValidade)
	}
}
