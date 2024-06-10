package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"fitness/domain"
)

var rorigemID string
var rdestinoID string
var rprodutoID int
var rquantidade int

var registrarRemessaCmd = &cobra.Command{
	Use:   "registrar-remessa",
	Short: "Registra uma remessa entre dois expositores",
	Long:  `Registra uma remessa entre dois expositores com os parâmetros fornecidos.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		guia, err := domain.NovoGuiaRemessa(rorigemID, rdestinoID, rprodutoID, rquantidade)
		if err != nil {
			fmt.Println("Erro ao criar guia de remessa:", err)
			return
		}

		origem := &domain.Expositor{
			ID: origemID,
			Estoque: map[int]int{
				produtoID: 100,
			},
		}
		destino := &domain.Expositor{
			ID:      destinoID,
			Estoque: map[int]int{},
		}

		remessa, err := guia.CriarRemessa(origem, destino)
		if err != nil {
			fmt.Println("Erro ao registrar remessa:", err)
			return
		}
		fmt.Printf("Remessa registrada:\n")
		fmt.Printf("OrigemID: %s\n", remessa.OrigemID)
		fmt.Printf("DestinoID: %s\n", remessa.DestinoID)
		fmt.Printf("ProdutoID: %d\n", remessa.ProdutoID)
		fmt.Printf("Quantidade: %d\n", remessa.Quantidade)
	},
}

func init() {
	rootCmd.AddCommand(registrarRemessaCmd)

	// Definindo flags para o comando
	registrarRemessaCmd.Flags().StringVarP(&rorigemID, "origemID", "o", "", "ID da origem (obrigatório)")
	registrarRemessaCmd.Flags().StringVarP(&rdestinoID, "destinoID", "d", "", "ID do destino (obrigatório)")
	registrarRemessaCmd.Flags().IntVarP(&rprodutoID, "produtoID", "p", 0, "ID do produto (obrigatório)")
	registrarRemessaCmd.Flags().IntVarP(&rquantidade, "quantidade", "q", 0, "Quantidade do produto (obrigatório)")

	// Marcando flags como obrigatórias
	registrarRemessaCmd.MarkFlagRequired("origemID")
	registrarRemessaCmd.MarkFlagRequired("destinoID")
	registrarRemessaCmd.MarkFlagRequired("produtoID")
	registrarRemessaCmd.MarkFlagRequired("quantidade")
}
