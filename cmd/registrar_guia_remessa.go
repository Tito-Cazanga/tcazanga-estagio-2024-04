package cmd

import (
	"fmt"
	"fitness/domain"
	"strconv"

	"github.com/spf13/cobra"
)

var registrarRemessaCmd = &cobra.Command{
	Use:   "registrar-remessa [origemID] [destinoID] [produtoID] [quantidade]",
	Short: "Registra uma remessa entre dois expositores",
	Long:  `Registra uma remessa entre dois expositores com os parâmetros fornecidos.`,
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

		origem := &domain.Expositor{
			ID: args[0],
			Estoque: map[int]int{
				produtoID: 100, // Exemplo de estoque inicial
			},
		}
		destino := &domain.Expositor{
			ID:      args[1],
			Estoque: map[int]int{},
		}

		remessa, err := guia.RegistrarRemessa(origem, destino)
		if err != nil {
			fmt.Println("Erro ao registrar remessa:", err)
			return
		}
		fmt.Printf("Remessa registrada: %+v\n", remessa)
	},
}

func init() {
	rootCmd.AddCommand(registrarRemessaCmd)
}
