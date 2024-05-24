package cmd

import (
	"fitness/application"
	"fitness/domain"
	"fmt"
	"strings"
	"strconv"

	"github.com/spf13/cobra"
)

var ginasioID string
var consumosFlag string

var emitirFaturaCmd = &cobra.Command{
	Use:   "emitir-fatura",
	Short: "Emite uma fatura para um ginásio com base nos consumos de produtos",
	Long:  `Emite uma fatura para um ginásio especificado com base nos consumos de produtos fornecidos.`,
	Args:  cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		consumos := []*domain.ConsumirProduto{}

		for _, consumoStr := range strings.Split(consumosFlag, ",") {
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
		fmt.Printf("Fatura emitida:\n")
		fmt.Printf("ID: %s\n", fatura.ID)
		fmt.Printf("Data de Emissão: %s\n", fatura.DataEmissao.Format("2006-01-02 15:04:05"))
		fmt.Printf("GinasioID: %s\n", fatura.GinasioID)
		fmt.Printf("Total: %.2f\n", fatura.Total)
		for _, consumo := range fatura.Consumos {
			fmt.Printf("Consumo - ExpositorID: %s, ProdutoID: %d, Quantidade: %d\n", consumo.ExpositorID, consumo.ProdutoID, consumo.Quantidade)
		}
	},
}

func init() {
	rootCmd.AddCommand(emitirFaturaCmd)

	// Definindo flags para o comando
	emitirFaturaCmd.Flags().StringVarP(&ginasioID, "ginasioID", "g", "", "ID do ginásio")
	emitirFaturaCmd.Flags().StringVarP(&consumosFlag, "consumos", "c", "", "Lista de consumos no formato expositorID:produtoID:quantidade separados por vírgula")

	// Marcando flags como obrigatórias
	emitirFaturaCmd.MarkFlagRequired("ginasioID")
	emitirFaturaCmd.MarkFlagRequired("consumos")
}
