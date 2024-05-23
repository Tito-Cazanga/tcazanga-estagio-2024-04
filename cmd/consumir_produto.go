package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"fitness/domain"
)

var consumirProdutoCmd = &cobra.Command{
	Use:   "consumir-produto [expositorID] [produtoID] [quantidade]",
	Short: "Consome um produto de um expositor",
	Long:  `Consome um produto do estoque de um expositor com os parâmetros fornecidos.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		produtoID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ProdutoID deve ser um número inteiro.")
			return
		}
		quantidade, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Quantidade deve ser um número inteiro.")
			return
		}
		
		consumirProduto, err := domain.NewConsumirProduto(args[0], produtoID, quantidade)
		if err != nil {
			fmt.Println("Erro ao criar comando de consumo:", err)
			return
		}

		expositor := &domain.Expositor{
			ID: args[0],
			Estoque: map[int]int{
				produtoID: 100, // Exemplo de estoque inicial
			},
		}

		consumo, err := consumirProduto.RegistrarConsumo(expositor)
		if err != nil {
			fmt.Println("Erro ao registrar consumo:", err)
			return
		}
		fmt.Printf("Consumo registrado: %+v\n", consumo)
	},
}

func init() {
	rootCmd.AddCommand(consumirProdutoCmd)
}
