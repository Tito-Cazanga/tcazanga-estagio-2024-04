package main

import (
	"fitness/application"
	"fmt"
)

func main(){
	expositor := &application.Expositor{
		ID:          "2",
		Localizacao: "Ginasio A",
		Estoque:     make(map[int]int),
	}

	comando := application.AbastecerExpositor{
		ExpositorID: "1",
		ProdutoID:   1,
		Quantidade:  10,
	}

	comando.AbastecerExpositor(expositor)
	for produtoID, quantidade := range expositor.Estoque {
		fmt.Printf("ProdutoID: %d, Quantidade: %d\n", produtoID, quantidade)
	}
	

}