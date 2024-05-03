package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Produto struct {
	ID        string
	Descricao string
	Categoria string
	Condicao  string
}

type Lote struct {
	ProdotoID     string
	DataExpiracao time.Time
	Quantidade    int
}
var produtos map[string]*Produto = make(map[string]*Produto)
var lotes map[string]*Lote = make(map[string]*Lote)

func main() {
	for {
		fmt.Println("\n1.Registrar Produto")
		fmt.Println("2. Registrar Lotes")
		fmt.Println("3. Monitorar Stock")
		fmt.Println("4. Sair")

		var escolha int
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			registroProduto()
		case 2:
			registroLote()
		case 3:
			monitorarlotes()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Opcao invalida, tenta de novo.")
		}
	}
}

func registroProduto() {
	var id, descricao, categoria, condicao string
	fmt.Print("Inserir ID Produto: ")
	fmt.Scan(&id)
	fmt.Print("Inserir Descricao do produto: ")
	fmt.Scan(&descricao)
	fmt.Print("Inserir Categoria do produto: ")
	fmt.Scan(&categoria)
	fmt.Print("Inserir condicao de armazenamento: ")
	fmt.Scan(&condicao)

	produto := &Produto{
		ID:        id,
		Descricao: descricao,
		Categoria: categoria,
		Condicao:  condicao,
	}

	produtos[id] = produto
	salvarProduto()
	fmt.Println("Produto registrado com sucesso.")
}

func registroLote() {
	var produtoID string
	var dataExpiracao time.Time
	var quantidade int

	fmt.Print("Inserir ID do lote: ")
	fmt.Scan(&produtoID)

	produto, exists := produtos[produtoID]
	if !exists {
		fmt.Println("Produto não encontrado.", produto)
		return
	}

	fmt.Print("Inserir a data de expiracao do lote (YYYY-MM-DD): ")
	fmt.Scan(&dataExpiracao)

	fmt.Print("Inserir da quanitdade de lote: ")
	fmt.Scan(&quantidade)

	lote := &Lote{
		ProdotoID:     produtoID,
		DataExpiracao: dataExpiracao,
		Quantidade:    quantidade,
	}

	lotes[produtoID] = lote
	salvarLotes()
	fmt.Println("Lote registrado com sucesso.")
}

func monitorarlotes() {
	for _, lote := range lotes {
		fmt.Printf("ID produto: %s, Data de expiracao: %s, Quantidade: %d\n", lote.ProdotoID, lote.DataExpiracao.Format("2024-01-02"), lote.Quantidade)
	}
}

func salvarProduto() {
	file, err := os.Create("produtos.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, produto := range produtos {
		record := []string{
			produto.ID,
			produto.Descricao,
			produto.Categoria,
			produto.Condicao,
		}
		writer.Write(record)
	}
}

func salvarLotes() {
	file, err := os.Create("lotes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, lote := range lotes {
		record := []string{
			lote.ProdotoID,
			lote.DataExpiracao.Format("2024-01-02"),
			fmt.Sprint(lote.Quantidade),
		}
		writer.Write(record)
	}
}
