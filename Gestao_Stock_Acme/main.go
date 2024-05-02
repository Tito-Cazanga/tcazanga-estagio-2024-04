package main

import (
	"fmt"
	"time"
)

type Produto struct {
	ID                    int
	Descricao             string
	Categoria             string
	CondicaoArmazenamento string
}

type Lote struct {
	Produto         *Produto
	NumeroDoLote    string
	DataDeEntrada   string
	DataDeExpiracao time.Time
	Quantidade      int
}

type Inventario struct {
	Produtos []*Produto
	Lotes    []*Lote
}

func NewInventario() *Inventario {
	return &Inventario{}
}

func (iv *Inventario) AdicionaProduto(produto *Produto) {
	iv.Produtos = append(iv.Produtos, produto)
}

func (iv *Inventario) AdicionaLote(lote *Lote) {
	iv.Lotes = append(iv.Lotes, lote)
}

func (inventario *Inventario) InserirNovoDado() {
	var id int
	var name, description, category, storageCondition string

	fmt.Print("Inserir nome do produto: ")
	fmt.Scan(&name)

	fmt.Print("Inserir a descrição do produto: ")
	fmt.Scan(&description)

	fmt.Print("Inserir categoria do produto: ")
	fmt.Scan(&category)

	fmt.Print("Inserir condicão de armazenamento do produto: ")
	fmt.Scan(&storageCondition)

	produto := inventario.Produtos

	id = len(produto) + 1

	novoProduto := Produto{
		ID:                    id,
		Descricao:             description,
		Categoria:             category,
		CondicaoArmazenamento: storageCondition,
	}

	iv := NewInventario()
	iv.AdicionaProduto(&novoProduto)

	fmt.Println("Produto adicionado com sucesso")
}

func 

func main(){

	println("====================================================")
	println("=========== SISTEMA DE GESTÃO DE STOCK ACME ======== ")
	println("====================================================")
	println("=============== ESCOLHA UMA DAS OPÇÕES  ============")
	println("=============== 1 - CADASTRAR PRODUTO ==============")
	println("=============== 2 - CADASTRAR PRODUTO EM UM LOTE ==========")
	println("=============== 3 - EXPORTAR INVENTÁRIO ==============")

	var opcao int
	fmt.Scan(&opcao)

	iv := NewInventario()

	switch opcao {
		case 1:
			iv.InserirNovoDado()
			
		default:
			fmt.Println("A opção selecionada não é válida.")	
	}

	
}
