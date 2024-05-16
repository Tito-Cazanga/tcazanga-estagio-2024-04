package application_test

import (
	application "fitness/application"
	"testing"
	"time"
)

func TestAbastecerExpositor(t *testing.T) {
	//Arrange
	expositor := &application.Expositor{
		ID:          "1",
		Localizacao: "Ginasio A",
		Estoque:     make(map[int]int),
	}

	comando := application.AbastecerExpositor{
		ExpositorID: "1",
		ProdutoID:   1,
		Quantidade:  10,
	}

	//Act
	evento := comando.AbastecerExpositor(expositor)

	//Assert
	if evento.ExpositorID != "1" || evento.ProdutoID != 1 || evento.Quantidade != 10 {
		t.Error("O expositor não foi abastecido")
	}

	if expositor.Estoque[1] != 10 {
		t.Error("O estoque do expositor está incorrecto após o abastecimento")
	}
}

func TestAbastecerExpositor_DoisAbastecimentos(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "2",
		Localizacao: "Ginásio D",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "2",
		ProdutoID:   1,
		Quantidade:  10,
	}

	// Act
	comando.AbastecerExpositor(expositor)
	evento2 := comando.AbastecerExpositor(expositor)

	// Assert
	if evento2.Quantidade != 10 {
		t.Errorf("O segundo abasteciemto deveria adicionar 10 unidades")
	}

	if expositor.Estoque[1] != 20 {
		t.Errorf("O estoque do expositor deveria ser 20 após o segundo abastecimento")
	}
}

func TestAbastecerExpositor_ProdutoDiferente(t *testing.T) {
	//Arrange
	expositor := &application.Expositor{
		ID:          "12",
		Localizacao: "Ginasio C",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "12",
		ProdutoID:   12,
		Quantidade:  20,
	}

	//Act
	evento := comando.AbastecerExpositor(expositor)

	//Assert
	if evento.ProdutoID != 12 {
		t.Error("ID do produto errado")
	}

	if expositor.Estoque[12] != 20 {
		t.Error("O estoque do expositor deveria ser 20 para produto 12")
	}
}

func TestAbastecerExpositor_ExpositorVazio(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "45",
		Localizacao: "Ginásio H",
		Estoque:     make(map[int]int),
	}

	comando := &application.AbastecerExpositor{
		ExpositorID: "45",
		ProdutoID:   14,
		Quantidade:  15,
	}

	// Act
	evento := comando.AbastecerExpositor(expositor)

	// Assert
	if evento.ExpositorID != "45" {
		t.Error("O evento deveria conter o ID do expositor 45")
	}

	if expositor.Estoque[14] != 15 {
		t.Errorf("O estoque do expositor deveria ser 15 para 0 produto %v", evento.ProdutoID)
	}
}

func TestVenderProduto_ProdutoNaoDisponivel(t *testing.T) {
	// Arrange: Definindo uma lista de produtos sem o produto a ser vendido
	produtos := []application.Produto{
		{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 10},
		{Nome: "Produto B", Validade: time.Now().Add(10 * 24 * time.Hour), Quantidade: 5},
	}

	// Act: Tentando vender um produto que não está disponível
	produtoVendido := "Produto C"
	quantidadeVendida := 3
	application.VenderProduto(&produtos, produtoVendido, quantidadeVendida)

	// Assert: Verificando se a quantidade disponível dos produtos permanece inalterada
	for _, p := range produtos {
		if p.Nome == produtoVendido {
			if p.Quantidade != 0 {
				t.Errorf("A quantidade disponível do produto não deveria ter sido alterada")
			}
		}
	}
}

func TestVenderProduto_ValidadesDiferentes(t *testing.T) {
	// Arrange: Definindo uma lista de produtos com o mesmo nome, mas com validades diferentes
	produtos := []application.Produto{
		{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 10},
		{Nome: "Produto A", Validade: time.Now().Add(10 * 24 * time.Hour), Quantidade: 5},
	}

	// Act: Vendendo o produto e verificando a quantidade disponível
	produtoVendido := "Produto A"
	quantidadeVendida := 3
	application.VenderProduto(&produtos, produtoVendido, quantidadeVendida)

	// Assert: Verificando se a quantidade disponível do produto com a validade mais antiga foi reduzida corretamente
	quantidadeEsperada := 7
	for _, p := range produtos {
		if p.Nome == produtoVendido && p.Validade == time.Now().Add(6*24*time.Hour) {
			if p.Quantidade != quantidadeEsperada {
				t.Errorf("A quantidade disponível do produto com validade mais antiga não foi reduzida corretamente")
			}
		}
	}
}

func TestVenderProduto_QuantidadeMaiorQueDisponivel(t *testing.T) {
	// Arrange
	produtos := []application.Produto{
		{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 80},
	}

	// Act
	produtoVendido := "Produto A"
	quantidadeVendida := 70 // Vendendo mais do que a disponível
	application.VenderProduto(&produtos, produtoVendido, quantidadeVendida)

	// Assert
	for _, p := range produtos {
		if p.Nome == produtoVendido {
			if p.Quantidade != 10 {
				t.Errorf("A quantidade disponível do produto deveria ser 0 após a venda exceder o estoque")
			}
		}
	}
}

func TestVendaDeProduto_AposExpositorAbastecido(t *testing.T) {
	// Arrange
	produtos := &[]application.Produto{
		{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 10},
		{Nome: "Produto B", Validade: time.Now().Add(10 * 24 * time.Hour), Quantidade: 5},
	}

	// Act
	produtoVendido := "Produto A"
	quantidadeVendida := 3
	application.VenderProduto(produtos, produtoVendido, quantidadeVendida)

	// Assert
	produtoEsperado := application.Produto{Nome: "Produto A", Validade: time.Now().Add(6 * 24 * time.Hour), Quantidade: 7}

	if !application.VerificarProdutoNoExpositor(produtos, produtoEsperado) {
		t.Errorf("Quantidade disponível do produto no expositor não foi reduzida corretamente após a venda")
	}
}

func TestConsumoRegistrado_AposExpositorAbastecido(t *testing.T) {
	// Arrange
	expositor := &application.Expositor{
		ID:          "1",
		Localizacao: "Ginásio A",
		Estoque:     map[int]int{1: 10},
	}
	comando := &application.ConsumirProduto{
		ExpositorID: 1,
		ProdutoID:   1,
		Quantidade:  5,
	}

	// Act
	evento := comando.ConsumirExpositor(expositor)

	// Assert
	if evento.GinásioID != "1" {
		t.Errorf("O ID do ginásio no evento está incorreto")
	}

	if evento.ProdutoID != 1 {
		t.Errorf("O ID do produto no evento está incorreto")
	}

	if evento.Quantidade != 5 {
		t.Errorf("A quantidade no evento está incorreta")
	}
}
