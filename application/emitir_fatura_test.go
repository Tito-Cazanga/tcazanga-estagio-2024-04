package application_test

import (
	"errors"
	"fitness/application"
	"fitness/domain"
	"testing"
)

func TestEmitirFatura(t *testing.T) {

	t.Run("Teste emitir uma nova fatura", func(t *testing.T) {
		// Arrange
		consumos := []*domain.ConsumoRegistrado{
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
			Consumos:  []*domain.ConsumoRegistrado{},
		}

		// Act
		fatura, err := comandoConsumosVazios.NewFatura()

		// Assert
		if !errors.Is(err, application.ErrListaConsumosVazia) {
			t.Errorf("Esperado: fatura vazia, recebido: %v", fatura)
		}
	})
}
