package domain_test

import (
	"fitness/domain"
	"testing"
	"time"
)

func TestEmitirFatura(t *testing.T) {

	consumos := []*domain.ConsumoRegistrado{
		{GinasioID: "ginásio1", ProdutoID: 1, Quantidade: 5},
		{GinasioID: "ginásio1", ProdutoID: 2, Quantidade: 10},
	}

	fatura, err := domain.FaturaEmitida("ginásio1", consumos)

	if err != nil {
		t.Errorf("Erro inesperado ao gerar a fatura: %v", err)
	}

	if fatura == nil {
		t.Error("A fatura gerada está nula")
	}

							
	t.Run("Emitir fatura consolidada", func(t *testing.T) {

		esperadoGinasioID := "ginásio1"   		
		if fatura.GinasioID != esperadoGinasioID {
			t.Errorf("ID do ginásio esperado: %s, obtido: %s", esperadoGinasioID, fatura.GinasioID)
		}
	
		now := time.Now()
		if !fatura.DataEmissao.After(now.Add(-time.Second)) || !fatura.DataEmissao.Before(now.Add(time.Second)) {
			t.Errorf("Data de emissão da fatura não está dentro da margem de erro: %v", fatura.DataEmissao)
		}
	
		esperadoTotal := 15.0 
		if fatura.Total != esperadoTotal {
			t.Errorf("Total da fatura esperado: %f, obtido: %f", esperadoTotal, fatura.Total)
		}
	
		esperadoNumConsumos := len(consumos)
		if len(fatura.Consumos) != esperadoNumConsumos {
			t.Errorf("Número de consumos na fatura esperado: %d, obtido: %d", esperadoNumConsumos, len(fatura.Consumos))
		}

	})
}
