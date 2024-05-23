package cmd

import (
	"fitness/application"
	"fitness/domain"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var ginasioID string
var consumosStr string

var emitCmd = &cobra.Command{
	Use:   "emit",
	Short: "Emite uma nova fatura",
	Run: func(cmd *cobra.Command, args []string) {
		consumos, err := parseConsumos(consumosStr)
		if err != nil {
			fmt.Printf("Erro ao analisar consumos: %v\n", err)
			return
		}

		comando := &application.EmitirFatura{
			GinasioID: ginasioID,
			Consumos:  consumos,
		}

		fatura, err := comando.NewFatura()
		if err != nil {
			fmt.Printf("Erro ao emitir fatura: %v\n", err)
			return
		}

		fmt.Printf("Fatura emitida com sucesso:\nID: %s\nData: %s\nTotal: %.2f\n", fatura.ID, fatura.DataEmissao, fatura.Total)
	},
}

func init() {
	emitCmd.Flags().StringVarP(&ginasioID, "ginasio", "g", "", "ID do ginásio (obrigatório)")
	emitCmd.Flags().StringVarP(&consumosStr, "consumos", "c", "", "Lista de consumos no formato 'produtoID:quantidade,produtoID:quantidade,...' (obrigatório)")
	emitCmd.MarkFlagRequired("ginasio")
	emitCmd.MarkFlagRequired("consumos")
}

func parseConsumos(consumosStr string) ([]*domain.ConsumirProduto, error) {
	var consumos []*domain.ConsumirProduto
	pares := strings.Split(consumosStr, ",")

	for _, par := range pares {
		parts := strings.Split(par, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("formato inválido para consumo: %s", par)
		}

		produtoID, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("produtoID inválido: %s", parts[0])
		}

		quantidade, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("quantidade inválida: %s", parts[1])
		}

		consumo := &domain.ConsumirProduto{
			GinasioID: ginasioID,
			ProdutoID: produtoID,
			Quantidade: quantidade,
		}
		consumos = append(consumos, consumo)
	}

	return consumos, nil
}
