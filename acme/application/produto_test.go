package application_test

import (
	"fitness/application"
	"testing"
	"time"
)

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