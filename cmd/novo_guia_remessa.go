package cmd

import (
	"fmt"
	"fitness/domain"
	"strconv"

	"github.com/spf13/cobra"
)

var novoGuiaRemessaCmd = &cobra.Command{
	Use:   "novo-guia-remessa [origemID] [destinoID] [produtoID] [quantidade]",
	Short: "Cria um novo Guia de Remessa",
	Long:  `Cria um novo Guia de Remessa com os parâmetros fornecidos.`,
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		produtoID, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("ProdutoID deve ser um número inteiro.")
			return
		}
		quantidade, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Quantidade deve ser um número inteiro.")
			return
		}
		guia, err := domain.NovoGuiaRemessa(args[0], args[1], produtoID, quantidade)
		if err != nil {
			fmt.Println("Erro ao criar guia de remessa:", err)
			return
		}
		fmt.Printf("Guia de Remessa criado: %+v\n", guia)
	},
}

func init() {
	rootCmd.AddCommand(novoGuiaRemessaCmd)
}
