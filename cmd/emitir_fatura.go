package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"fitness/application"
	"fitness/domain"
)

var emitirFaturaCmd = &cobra.Command{
	Use:   "emitir-fatura [ginasioID] [consumos]",
	Short: "Emite uma fatura para um ginásio com base nos consumos de produtos",
	Long:  `Emite uma fatura para um ginásio especificado com base nos consumos de produtos fornecidos.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ginasioID := args[0]
		consumos := []*domain.ConsumirProduto{}

		for _, consumoStr := range strings.Split(args[1], ",") {
			parts := strings.Split(consumoStr, ":")
			if len(parts) != 3 {
				fmt.Println("Formato de consumo inválido. Use: expositorID:produtoID:quantidade")
				return
			}

			produtoID, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("ProdutoID deve ser um número inteiro.")
				return
			}
			quantidade, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Quantidade deve ser um número inteiro.")
				return
			}

			consumos = append(consumos, &domain.ConsumirProduto{
				ExpositorID: parts[0],
				ProdutoID:   produtoID,
				Quantidade:  quantidade,
			})
		}

		emitirFaturaCmd := &application.EmitirFatura{
			GinasioID: ginasioID,
			Consumos:  consumos,
		}

		fatura, err := emitirFaturaCmd.NewFatura()
		if err != nil {
			fmt.Println("Erro ao emitir fatura:", err)
			return
		}
		fmt.Printf("Fatura emitida: %+v\n", fatura)
	},
}

func init() {
	rootCmd.AddCommand(emitirFaturaCmd)
}
