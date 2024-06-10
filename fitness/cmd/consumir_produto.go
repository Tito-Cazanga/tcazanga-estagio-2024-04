package cmd

import (
	"fitness/domain"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var consumirProdutoCmd = &cobra.Command{
	Use:   "consumir-produto",
	Short: "Consome um produto de um expositor",
	Long:  `Consome um produto do estoque de um expositor com os parâmetros fornecidos.`,
	Args:  cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		expositorID, _ := cmd.Flags().GetString("expositorID")
		produtoID, _ := cmd.Flags().GetInt("produtoID")
		quantidade, _ := cmd.Flags().GetInt("quantidade")

		consumirProduto, err := domain.NewConsumirProduto(expositorID, produtoID, quantidade)
		if err != nil {
			fmt.Println("Erro ao criar comando de consumo:", err)
			return
		}

		expositor := &domain.Expositor{
			ID: expositorID,
			Estoque: map[int]int{
				produtoID: 100, 
			},
		}

		consumo, err := consumirProduto.RegistrarConsumo(expositor)
		if err != nil {
			fmt.Println("Erro ao registrar consumo:", err)
			return
		}
		fmt.Printf("Consumo registrado:\n")
		fmt.Printf("ExpositorID: %s\n", consumo.ExpositorID)
		fmt.Printf("ProdutoID: %d\n", consumo.ProdutoID)
		fmt.Printf("Quantidade: %d\n", consumo.Quantidade)
		fmt.Printf("Data: %s\n", time.Now().Format(time.RFC3339))
	},
}

func init() {
	rootCmd.AddCommand(consumirProdutoCmd)

	// Definindo flags para o comando
	consumirProdutoCmd.Flags().StringVarP(&expositorID, "expositorID", "e", "", "ID do expositor")
	consumirProdutoCmd.Flags().IntVarP(&produtoID, "produtoID", "p", 0, "ID do produto")
	consumirProdutoCmd.Flags().IntVarP(&quantidade, "quantidade", "q", 0, "Quantidade de produtos")

	// Marcando flags como obrigatórias
	consumirProdutoCmd.MarkFlagRequired("expositorID")
	consumirProdutoCmd.MarkFlagRequired("produtoID")
	consumirProdutoCmd.MarkFlagRequired("quantidade")
}
