package cmd

import (
	"fitness/domain"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var abastecerExpositorCmd = &cobra.Command{
	Use:   "abastecer-expositor [expositorID] [produtoID] [quantidade]",
	Short: "Abastece um expositor com um determinado produto",
	Long:  `Abastece um expositor com um determinado produto, respeitando a restrição de dois abastecimentos por semana.`,
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

		expositor := &domain.Expositor{
			ID: args[0],
			Estoque: map[int]int{
				produtoID: 50, 
			},
			Abastecimentos: []time.Time{}, 
		}

		abastecerCmd := &domain.AbastecerExpositor{
			ExpositorID: args[0],
			ProdutoID:   produtoID,
			Quantidade:  quantidade,
		}

		abastecimento, err := abastecerCmd.AbastecerExpositor(expositor)
		if err != nil {
			fmt.Println("Erro ao abastecer expositor:", err)
			return
		}
		fmt.Printf("Abastecimento realizado: %+v\n", abastecimento)
	},
}

func init() {
	rootCmd.AddCommand(abastecerExpositorCmd)
}
