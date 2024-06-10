package domain_test

import (
	"testing"
	"time"

	"fitness/domain"
)

func TestAbastecerExpositor(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:          "1",
		Localizacao: "Ginasio A",
		Estoque:     make(map[int]int),
	}

	comando := &domain.AbastecerExpositor{
		ExpositorID: "1",
		ProdutoID:   1,
		Quantidade:  10,
	}

	// Act
	evento := comando.Abastecer(expositor)

	// Assert
	if evento.ExpositorID != "1" || evento.ProdutoID != 1 || evento.Quantidade != 10 {
		t.Error("O evento retornado pelo comando está incorreto")
	}

	if expositor.Estoque[1] != 10 {
		t.Error("O estoque do expositor está incorreto após o abastecimento")
	}
}

func TestAbastecerExpositor_DoisAbastecimentos(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:          "2",
		Localizacao: "Ginásio D",
		Estoque:     make(map[int]int),
	}

	comando := &domain.AbastecerExpositor{
		ExpositorID: "2",
		ProdutoID:   1,
		Quantidade:  10,
	}

	// Act
	comando.Abastecer(expositor)
	evento2 := comando.Abastecer(expositor)

	// Assert
	if evento2.Quantidade != 10 {
		t.Errorf("O segundo abastecimento deveria adicionar 10 unidades")
	}

	if expositor.Estoque[1] != 20 {
		t.Errorf("O estoque do expositor deveria ser 20 após o segundo abastecimento")
	}
}

func TestAbastecerExpositor_ProdutoDiferente(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:          "12",
		Localizacao: "Ginasio C",
		Estoque:     make(map[int]int),
	}

	comando := &domain.AbastecerExpositor{
		ExpositorID: "12",
		ProdutoID:   12,
		Quantidade:  20,
	}

	// Act
	evento := comando.Abastecer(expositor)

	// Assert
	if evento.ProdutoID != 12 {
		t.Error("ID do produto errado no evento")
	}

	if expositor.Estoque[12] != 20 {
		t.Error("O estoque do expositor deveria ser 20 para o produto 12")
	}
}

func TestAbastecerExpositor_ExpositorVazio(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:          "45",
		Localizacao: "Ginásio H",
		Estoque:     make(map[int]int),
	}

	comando := &domain.AbastecerExpositor{
		ExpositorID: "45",
		ProdutoID:   14,
		Quantidade:  15,
	}

	// Act
	evento := comando.Abastecer(expositor)

	// Assert
	if evento.ExpositorID != "45" {
		t.Error("O evento deveria conter o ID do expositor 45")
	}

	if expositor.Estoque[14] != 15 {
		t.Errorf("O estoque do expositor deveria ser 15 para o produto %v", evento.ProdutoID)
	}
}

func TestAbastecerExpositor_DoisAbastecimentosPorSemana(t *testing.T) {
	// Arrange
	expositor := &domain.Expositor{
		ID:             "2",
		Localizacao:    "Ginásio D",
		Estoque:        make(map[int]int),
		Abastecimentos: []time.Time{},
	}

	comando := &domain.AbastecerExpositor{
		ExpositorID: "2",
		ProdutoID:   1,
		Quantidade:  10,
	}

	// Act 
	_, err := comando.AbastecerExpositor(expositor)
	if err != nil {
		t.Fatalf("Erro ao abastecer pela primeira vez: %v", err)
	}

	// Act 
	_, err = comando.AbastecerExpositor(expositor)
	if err != nil {
		t.Fatalf("Erro ao abastecer pela segunda vez: %v", err)
	}

	// Act
	_, err = comando.AbastecerExpositor(expositor)
	if err == nil {
		t.Fatal("Esperava-se erro para terceiro abastecimento na mesma semana, mas não ocorreu")
	}

	// Assert
	if expositor.Estoque[1] != 20 {
		t.Errorf("O estoque do expositor deveria ser 20 após dois abastecimentos. Obtido: %d", expositor.Estoque[1])
	}

	if len(expositor.Abastecimentos) != 2 {
		t.Errorf("O número de abastecimentos registrados deveria ser 2. Obtido: %d", len(expositor.Abastecimentos))
	}
}
