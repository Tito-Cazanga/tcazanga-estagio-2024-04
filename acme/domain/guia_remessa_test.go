package domain_test

import (
	"testing"

	"fitness/domain"
)

func TestRegistrarRemessa(t *testing.T) {
	// Arrange
	origem := &domain.Expositor{
		ID:          "CentroDistribuicao",
		Localizacao: "Central",
		Estoque:     map[int]int{1: 50},
	}
	destino := &domain.Expositor{
		ID:          "GinásioA",
		Localizacao: "Bairro A",
		Estoque:     map[int]int{1: 50},
	}

	comando, err := domain.NovoGuiaRemessa("CentroDistribuicao", "GinásioA", 1, 50)
	if err != nil {
		t.Fatalf("Erro ao criar GuiaRemessa: %v", err)
	}

	// Act
	evento, err := comando.CriarRemessa(origem, destino)
	if err != nil {
		t.Fatalf("Erro ao registrar remessa: %v", err)
	}

	// Assert
	if evento.OrigemID != "CentroDistribuicao" {
		t.Errorf("O ID da origem no evento está incorreto. Esperado: 'CentroDistribuicao', Obtido: '%s'", evento.OrigemID)
	}

	if evento.DestinoID != "GinásioA" {
		t.Errorf("O ID do destino no evento está incorreto. Esperado: 'GinásioA', Obtido: '%s'", evento.DestinoID)
	}

	if evento.ProdutoID != 1 {
		t.Errorf("O ID do produto no evento está incorreto. Esperado: '1', Obtido: '%d'", evento.ProdutoID)
	}

	if evento.Quantidade != 50 {
		t.Errorf("A quantidade no evento está incorreta. Esperado: '50', Obtido: '%d'", evento.Quantidade)
	}

	if origem.Estoque[1] != 50 {
		t.Errorf("O estoque da origem está incorreto após a remessa. Esperado: '50', Obtido: '%d'", origem.Estoque[1])
	}

	if destino.Estoque[1] != 50 {
		t.Errorf("O estoque do destino está incorreto após a remessa. Esperado: '50', Obtido: '%d'", destino.Estoque[1])
	}
}

func TestRegistrarRemessa_EstoqueInsuficiente(t *testing.T) {
	// Arrange
	origem := &domain.Expositor{
		ID:          "CentroDistribuicao",
		Localizacao: "Central",
		Estoque:     map[int]int{1: 30},
	}
	destino := &domain.Expositor{
		ID:          "GinásioA",
		Localizacao: "Bairro A",
		Estoque:     map[int]int{},
	}

	comando, err := domain.NovoGuiaRemessa("CentroDistribuicao", "GinásioA", 1, 50)
	if err != nil {
		t.Fatalf("Erro ao criar GuiaRemessa: %v", err)
	}

	// Act
	evento, err := comando.CriarRemessa(origem, destino)

	// Assert
	if err == nil {
		t.Fatal("Esperava-se um erro de estoque insuficiente, mas não houve erro")
	}

	if evento != nil {
		t.Fatal("Evento não deveria ser gerado em caso de erro")
	}

	if origem.Estoque[1] != 30 {
		t.Errorf("O estoque da origem não deveria mudar em caso de erro. Esperado: '30', Obtido: '%d'", origem.Estoque[1])
	}

	if destino.Estoque[1] != 0 {
		t.Errorf("O estoque do destino não deveria mudar em caso de erro. Esperado: '0', Obtido: '%d'", destino.Estoque[1])
	}
}
