package application_test

import (
	"errors"
	"fitness/application"
	"fitness/domain"
	"testing"
	"time"
)

func TestEmitirFatura1(t *testing.T) {

	t.Run("Teste emitir uma nova fatura", func(t *testing.T) {
		// Arrange
		consumos := []*domain.ConsumirProduto{
			{GinasioID: "ginásio1", ProdutoID: 1, Quantidade: 5},
			{GinasioID: "ginásio1", ProdutoID: 2, Quantidade: 10},
		}

		comando := &application.EmitirFatura{
			GinasioID: "ginásio1",
			Consumos:  consumos,
		}

		// Act
		fatura, err := comando.NewFatura()

		// Assert
		if err != nil {
			t.Errorf("Erro inesperado ao emitir a fatura: %v", err)
		}

		if fatura == nil {
			t.Error("A fatura emitida está nula")
		}

		if fatura.GinasioID != "ginásio1" {
			t.Errorf("ID do ginásio na fatura incorreto: esperado: ginásio1, recebido: %s", fatura.GinasioID)
		}
	})

	t.Run("Teste emitir uma nova mesmo sem consumos no ginásio", func(t *testing.T) {
		// Arrange
		comandoConsumosVazios := &application.EmitirFatura{
			GinasioID: "ginásio2",
			Consumos:  []*domain.ConsumirProduto{},
		}

		// Act
		fatura, err := comandoConsumosVazios.NewFatura()

		// Assert
		if !errors.Is(err, application.ErrListaConsumosVazia) {
			t.Errorf("Esperado: fatura vazia, recebido: %v", fatura)
		}
	})
}

func TestEmitirFatura(t *testing.T) {

	consumos := []*domain.ConsumirProduto{
		{GinasioID: "ginásio1", ProdutoID: 1, Quantidade: 5},
		{GinasioID: "ginásio1", ProdutoID: 2, Quantidade: 10},
	}

	fatura, err := application.FaturaEmitida("ginásio1", consumos)

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

