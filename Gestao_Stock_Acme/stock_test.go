package stock_test

import (
	"stock"
	"testing"
)

func teste_Produto(t *testing.T, produto[]string, resultado[] string){
	t.Fail()
}

func Erro(t *testing.T){
	//Arrange
	add := stock.Produto{}

	//act
	add.AdicionarProduto()
	
	//Assert
	 teste_Produto(t,stock.Produto{ID: 13,
		Categoria: "Bebida",
		Detalhe: "Sei la",
		EstadoProduto: "ambiental",
		Validade: "12",
		},)
}