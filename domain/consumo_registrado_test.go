package domain_test

import (
	"fitness/domain"
	"testing"
)

func TestNewConsumoRegistrado(t *testing.T) {
	
	ginásioID := "ginásio1"
	produtoID := 1
	quantidade := 10

	consumo, err :=  domain.NewConsumoRegistrado(ginásioID, produtoID, quantidade)

	t.Run("Quantidade positiva", func(t *testing.T) {
		
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
	
		if consumo == nil {
			t.Error("O objeto ConsumoRegistrado não deveria ser nulo")
		}
	
		if consumo.GinásioID != ginásioID {
			t.Errorf("GinásioID esperado: %s, obtido: %s", ginásioID, consumo.GinásioID)
		}
	
		if consumo.ProdutoID != produtoID {
			t.Errorf("ProdutoID esperado: %d, obtido: %d", produtoID, consumo.ProdutoID)
		}
	
		if consumo.Quantidade != quantidade {
			t.Errorf("Quantidade esperada: %d, obtida: %d", quantidade, consumo.Quantidade)
		}
	})

	t.Run("Quantidade não positiva", func(t *testing.T) {
		quantidadeNegativa := -5

		consumo, err = domain.NewConsumoRegistrado(ginásioID, produtoID, quantidadeNegativa)
	
		if err == nil {
			t.Error("Esperava-se um erro para quantidade não positiva")
		}
	
		if consumo != nil {
			t.Error("O objeto ConsumoRegistrado deveria ser nulo")
		}

	})
	
}
