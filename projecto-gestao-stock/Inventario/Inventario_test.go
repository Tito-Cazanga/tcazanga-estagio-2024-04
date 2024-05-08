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

func TestNovoProdutoNome(t *testing.T) {
    
    dataDeValidade := time.Now().AddDate(0, 1, 0) 

    p := inventario.NovoProduto(1, "Produto Teste", dataDeValidade)

    if p.Nome != "Produto Teste" {
        t.Errorf("Nome esperado: %s, Nome retornado: %s", "Produto Teste", p.Nome)
    }
}

func TestNovoProdutoDataDeValidade(t *testing.T) {
    dataDeValidadeEsperada := time.Now().AddDate(2024, 05, 8) 
    
    p := inventario.NovoProduto(1, "Produto Teste", dataDeValidadeEsperada)

    if p.DataDeValidade != dataDeValidadeEsperada {
        t.Errorf("Data de validade esperada: %s, Data de validade retornada: %s", dataDeValidadeEsperada, p.DataDeValidade)
    }
}

func TestNovoLoteNaoNulo(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    if lote == nil {
        t.Error("O lote não deve ser nulo.")
    }
}

func TestNovoLoteIDDoProduto(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    if lote.ProdutoID != 1 {
        t.Errorf("ID do produto esperado: %d, ID do produto retornado: %d", 1, lote.ProdutoID)
    }
}

func TestNovoLoteLote(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    if lote.Lote != "Lote001" {
        t.Errorf("Lote esperado: %s, Lote retornado: %s", "Lote001", lote.Lote)
    }
}

func TestNovoLoteQuantidade(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    if lote.Quantidade != 100 {
        t.Errorf("Quantidade esperada: %d, Quantidade retornada: %d", 100, lote.Quantidade)
    }
}