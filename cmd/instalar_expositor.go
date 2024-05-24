package cmd

import (
	"fitness/domain"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var instalarExpositorCmd = &cobra.Command{
	Use:   "instalar-expositor [expositorID] [localizacao] [produtos]",
	Short: "Instala um novo expositor com produtos iniciais",
	Long:  `Instala um novo expositor em uma localização específica com produtos iniciais.`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		produtos := make(map[int]int)
		// Parse produtos from the argument
		// Assumes format: produtoID1:quantidade1,produtoID2:quantidade2,...
		for _, produto := range strings.Split(args[2], ",") {
			parts := strings.Split(produto, ":")
			produtoID, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("ProdutoID deve ser um número inteiro.")
				return
			}
			quantidade, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Quantidade deve ser um número inteiro.")
				return
			}
			produtos[produtoID] = quantidade
		}

		expositorInstalado := domain.NovoExpositorInstalado(args[0], args[1], produtos)
		fmt.Printf("Expositor instalado: %+v\n", expositorInstalado)
	},
}

func init() {
	rootCmd.AddCommand(instalarExpositorCmd)
}
