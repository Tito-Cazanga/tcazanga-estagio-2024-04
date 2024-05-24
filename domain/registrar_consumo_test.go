package domain_test

import (
	"fitness/domain"
	"testing"
)

func TestConsumoRegistrado_AposExpositorAbastecido(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:          "1",
		Localizacao: "Ginásio A",
		Estoque:     map[int]int{1: 10},
	}
	comando := &domain.ConsumirProduto{
		GinasioID: "1",
		ProdutoID:   1,
		Quantidade:  5,
	}

	// Act
	evento := comando.RegistrarConsumo(expositor)

	// Assert
	if evento.GinasioID != "1" {
		t.Errorf("O ID do ginásio no evento está incorreto")
	}

	if evento.ProdutoID != 1 {
		t.Errorf("O ID do produto no evento está incorreto")
	}

	if evento.Quantidade != 5 {
		t.Errorf("A quantidade no evento está incorreta")
	}
	
}


func TestNewConsumoRegistrado(t *testing.T) {
	
	ginasioID := "ginásio1"
	produtoID := 1
	quantidade := 10

	consumo, err :=  domain.NewConsumirProduto(ginasioID, produtoID, quantidade)

	t.Run("Quantidade positiva", func(t *testing.T) {
		
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
	
		if consumo == nil {
			t.Error("O objeto ConsumoRegistrado não deveria ser nulo")
		}
	
		if consumo.GinasioID != ginasioID {
			t.Errorf("GinásioID esperado: %s, obtido: %s", ginasioID, consumo.GinasioID)
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

		consumo, err = domain.NewConsumirProduto(ginasioID, produtoID, quantidadeNegativa)
	
		if err == nil {
			t.Error("Esperava-se um erro para quantidade não positiva")
		}
	
		if consumo != nil {
			t.Error("O objeto ConsumoRegistrado deveria ser nulo")
		}

	})
	
}
