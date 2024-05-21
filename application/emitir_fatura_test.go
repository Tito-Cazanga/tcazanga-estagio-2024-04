package application_test

import (
	"errors"
	"fitness/application"
	"fitness/domain"
	"testing"
)

func TestEmitirFatura_Execute(t *testing.T) {
	consumos := []*domain.ConsumoRegistrado{
		{GinasioID: "ginásio1", ProdutoID: 1, Quantidade: 5},
		{GinasioID: "ginásio1", ProdutoID: 2, Quantidade: 10},
	}

	comando := &application.EmitirFatura{
		GinasioID: "ginásio1",
		Consumos:  consumos,
	}

	fatura, err := comando.Executar()
	t.Run("Consumos não vazios", func(t *testing.T) {

		if err != nil {
			t.Errorf("Erro inesperado ao emitir a fatura: %v", err)
		}

		if fatura == nil {
			t.Error("A fatura emitida está nula")
		}

		if fatura.GinasioID != "ginásio1" {
			t.Errorf("ID do ginásio na fatura incorreto: esperado: ginásio1, obtido: %s", fatura.GinasioID)
		}
	})

	t.Run("Consumos vazios", func(t *testing.T) {
		comandoConsumosVazios := &application.EmitirFatura{
			GinasioID: "ginásio2",
			Consumos:  []*domain.ConsumoRegistrado{},
		}

		fatura, err = comandoConsumosVazios.Executar()

		if err == nil {
			t.Error("Esperava-se um erro ao emitir uma fatura com consumos vazios")
		}

		if !errors.Is(err, application.ErrListaConsumosVazia) {
			t.Errorf("Erro esperado ErrListaConsumosVazia não encontrado: %v", err)
		}
	})
}
