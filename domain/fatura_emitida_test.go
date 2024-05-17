package domain

import (
	"testing"
	"time"
)

func TestEmitirFatura(t *testing.T) {

	consumos := []*ConsumoRegistrado{
		{GinásioID: "ginásio1", ProdutoID: 1, Quantidade: 5},
		{GinásioID: "ginásio1", ProdutoID: 2, Quantidade: 10},
	}

	fatura, err := EmitirFatura("ginásio1", consumos)


	if err != nil {
		t.Errorf("Erro inesperado ao gerar a fatura: %v", err)
	}

	if fatura == nil {
		t.Error("A fatura gerada está nula")
	}

	expectedGinásioID := "ginásio1"
	if fatura.GinásioID != expectedGinásioID {
		t.Errorf("ID do ginásio esperado: %s, obtido: %s", expectedGinásioID, fatura.GinásioID)
	}

	now := time.Now()
	if !fatura.DataEmissao.After(now.Add(-time.Second)) || !fatura.DataEmissao.Before(now.Add(time.Second)) {
		t.Errorf("Data de emissão da fatura não está dentro da margem de erro: %v", fatura.DataEmissao)
	}

	expectedTotal := 15.0 
	if fatura.Total != expectedTotal {
		t.Errorf("Total da fatura esperado: %f, obtido: %f", expectedTotal, fatura.Total)
	}

	expectedNumConsumos := len(consumos)
	if len(fatura.Consumos) != expectedNumConsumos {
		t.Errorf("Número de consumos na fatura esperado: %d, obtido: %d", expectedNumConsumos, len(fatura.Consumos))
	}
}
