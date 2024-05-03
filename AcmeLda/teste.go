package main

import "testing"

func TestProduto(t *testing.T) {
	// Arrange: Preparar o estado do teste
	esperadoProduto := Produto{
		ID:        "123",
		Descricao: "Produto de Teste",
		Categoria: "Categoria de Teste",
		Condicao:  "Boa",
	}

	// Act: Executar a operação que será testada
	produto := Produto{
		ID:        "123",
		Descricao: "Produto de Teste",
		Categoria: "Categoria de Teste",
		Condicao:  "Boa",
	}

	// Assert: Verificar se o resultado é o esperado
	if produto != esperadoProduto {
		t.Errorf("O produto esperado %+v é diferente do produto obtido %+v", esperadoProduto, produto)
	}
}
