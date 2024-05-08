package main_test

import (
	inventario "AcmeStock/Inventario"
	"testing"
	"time"
)

func TestNovoProdutoNaoNulo(t *testing.T) { 
    p := inventario.NovoProduto(1, "Produto Teste")

    if p == nil {
        t.Error("O produto não deve ser nulo.")
    }
}

func TestNovoProdutoID(t *testing.T) {
    
    p := inventario.NovoProduto(1, "Produto Teste")

    if p.ID != 1 {
        t.Errorf("ID esperado: %d, ID retornado: %d", 1, p.ID)
    }
}

func TestNovoProdutoNome(t *testing.T) {

    p := inventario.NovoProduto(1, "Produto Teste")

    if p.Nome != "Produto Teste" {
        t.Errorf("Nome esperado: %s, Nome retornado: %s", "Produto Teste", p.Nome)
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

func TestNovoLoteLocalizacao(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    corredorEsperado := "CorredorA"
    armarioEsperado := "Armario1"
    prateleiraEsperada := "Prateleira2"

    if lote.Localizacao.Corredor != corredorEsperado || lote.Localizacao.Armario != armarioEsperado || lote.Localizacao.Prateleira != prateleiraEsperada {
        t.Errorf("Localização esperada: Corredor %s, Armário %s, Prateleira %s, Localização retornada: Corredor %s, Armário %s, Prateleira %s", corredorEsperado, armarioEsperado, prateleiraEsperada, lote.Localizacao.Corredor, lote.Localizacao.Armario, lote.Localizacao.Prateleira)
    }
}

func TestNovoLoteDataDeValidade(t *testing.T) {
    lote := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)
    
    dataDeValidadeEsperada := time.Date(2024, time.Month(6), 1, 0, 0, 0, 0, time.UTC)

    if lote.DataDeValidade != dataDeValidadeEsperada {
        t.Errorf("Data de validade esperada: %s, Data de validade retornada: %s", dataDeValidadeEsperada, lote.DataDeValidade)
    }
}

func TestRecomendarLotes(t *testing.T) {

    lote1 := inventario.NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)

    lote2 := inventario.NovoLote(2, "Lote002", 50, "CorredorB", "Armario2", "Prateleira1", 2024, 6, 15)

    lote3 := inventario.NovoLote(3, "Lote003", 200, "CorredorC", "Armario3", "Prateleira3", 2024, 5, 20)

    lotes := []*inventario.Lote{lote1, lote2, lote3}

    lotesRecomendados := inventario.RecomendarLotes(lotes)

    if len(lotesRecomendados) != len(lotes) {
        t.Errorf("O número de lotes recomendados (%d) deve ser igual ao número de lotes originais (%d)", len(lotesRecomendados), len(lotes))
    }

    // Verifica se os lotes estão ordenados por proximidade da validade
    for i := 0; i < len(lotesRecomendados)-1; i++ {

        if lotesRecomendados[i].DiasParaValidade() > lotesRecomendados[i+1].DiasParaValidade() {
            t.Errorf("Os lotes não estão ordenados corretamente por proximidade da validade")
            break
        }
    }
}