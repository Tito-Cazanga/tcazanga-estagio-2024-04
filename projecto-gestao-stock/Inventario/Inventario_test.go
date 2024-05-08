package inventario_test

import (
	inventario "AcmeStock/Inventario"
	"testing"
	"time"
)

func TestNovoProdutoNaoNulo(t *testing.T) {
    dataDeValidade := time.Now().AddDate(0, 1, 0) // Adiciona um mês à data atual
    p := inventario.NovoProduto(1, "Produto Teste", dataDeValidade)

    if p == nil {
        t.Error("O produto não deve ser nulo.")
    }
}
