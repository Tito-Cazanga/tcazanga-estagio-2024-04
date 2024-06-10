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
	comando, err := domain.NewConsumirProduto("1", 1, 5)
	if err != nil {
		t.Fatalf("Erro ao criar comando ConsumirProduto: %v", err)
	}

	// Act
	evento, err := comando.RegistrarConsumo(expositor)
	if err != nil {
		t.Fatalf("Erro ao registrar consumo: %v", err)
	}

	// Assert
	if evento.ExpositorID != "1" {
		t.Errorf("O ID do expositor no evento está incorreto. Esperado: '1', Obtido: '%s'", evento.ExpositorID)
	}

	if evento.ProdutoID != 1 {
		t.Errorf("O ID do produto no evento está incorreto. Esperado: '1', Obtido: '%d'", evento.ProdutoID)
	}

	if evento.Quantidade != 5 {
		t.Errorf("A quantidade no evento está incorreta. Esperado: '5', Obtido: '%d'", evento.Quantidade)
	}

	if expositor.Estoque[1] != 5 {
		t.Errorf("O estoque do expositor está incorreto após o consumo. Esperado: '5', Obtido: '%d'", expositor.Estoque[1])
	}

}

func TestNewConsumirProduto(t *testing.T) {
	ginasioID := "ginásio1"
	produtoID := 1
	quantidade := 10

	t.Run("Quantidade positiva", func(t *testing.T) {
		consumo, err := domain.NewConsumirProduto(ginasioID, produtoID, quantidade)

		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		if consumo == nil {
			t.Error("O objeto ConsumirProduto não deveria ser nulo")
		}

		if consumo.ExpositorID != ginasioID {
			t.Errorf("GinásioID esperado: %s, obtido: %s", ginasioID, consumo.ExpositorID)
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

		consumo, err := domain.NewConsumirProduto(ginasioID, produtoID, quantidadeNegativa)

		if err == nil {
			t.Error("Esperava-se um erro para quantidade não positiva")
		}

		if consumo != nil {
			t.Error("O objeto ConsumirProduto deveria ser nulo")
		}
	})
}
