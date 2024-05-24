package cmd

import (
	"fitness/domain"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var produtoID int
var quantidade int
var expositorID string

var abastecerExpositorCmd = &cobra.Command{
	Use:   "abastecer-expositor",
	Short: "Abastece um expositor com um determinado produto",
	Long:  `Abastece um expositor com um determinado produto, respeitando a restrição de dois abastecimentos por semana.`,
	Args:  cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		expositorID, _ := cmd.Flags().GetString("expositorID")
		produtoID, _ := cmd.Flags().GetInt("produtoID")
		quantidade, _ := cmd.Flags().GetInt("quantidade")

		expositor := &domain.Expositor{
			ID: expositorID,
			Estoque: map[int]int{
				produtoID: 50, 
			},
			Abastecimentos: []time.Time{}, 
		}

		abastecerCmd := &domain.AbastecerExpositor{
			ExpositorID: expositorID,
			ProdutoID:   produtoID,
			Quantidade:  quantidade,
		}

		abastecimento, err := abastecerCmd.AbastecerExpositor(expositor)
		if err != nil {
			fmt.Println("Erro ao abastecer expositor:", err)
			return
		}
		fmt.Printf("Abastecimento realizado:\n")
		fmt.Printf("ExpositorID: %s\n", abastecimento.ExpositorID)
		fmt.Printf("ProdutoID: %d\n", abastecimento.ProdutoID)
		fmt.Printf("Quantidade: %d\n", abastecimento.Quantidade)
		fmt.Printf("Data: %s\n", time.Now().Format(time.RFC3339))
	},
}

func init() {
	rootCmd.AddCommand(abastecerExpositorCmd)

	// Definindo flags para o comando
	abastecerExpositorCmd.Flags().StringVarP(&expositorID, "expositorID", "e", "", "ID do expositor")
	abastecerExpositorCmd.Flags().IntVarP(&produtoID, "produtoID", "p", 0, "ID do produto")
	abastecerExpositorCmd.Flags().IntVarP(&quantidade, "quantidade", "q", 0, "Quantidade de produtos")

	// Marcando flags como obrigatórias
	abastecerExpositorCmd.MarkFlagRequired("expositorID")
	abastecerExpositorCmd.MarkFlagRequired("produtoID")
	abastecerExpositorCmd.MarkFlagRequired("quantidade")
}
