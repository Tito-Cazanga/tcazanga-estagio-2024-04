package cmd

import (
	"fitness/domain"
	"fmt"

	"github.com/spf13/cobra"
)

var origemID string
var destinoID string
var gprodutoID int
var gquantidade int

var novoGuiaRemessaCmd = &cobra.Command{
	Use:   "novo-guia-remessa",
	Short: "Cria um novo Guia de Remessa",
	Long:  `Cria um novo Guia de Remessa com os parâmetros fornecidos.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		guia, err := domain.NovoGuiaRemessa(origemID, destinoID, gprodutoID, gquantidade)
		if err != nil {
			fmt.Println("Erro ao criar guia de remessa:", err)
			return
		}
		fmt.Printf("Guia de Remessa criado:\n")
		fmt.Printf("OrigemID: %s\n", guia.OrigemID)
		fmt.Printf("DestinoID: %s\n", guia.DestinoID)
		fmt.Printf("ProdutoID: %d\n", guia.ProdutoID)
		fmt.Printf("Quantidade: %d\n", guia.Quantidade)
	},
}

func init() {
	rootCmd.AddCommand(novoGuiaRemessaCmd)

	// Definindo flags para o comando
	novoGuiaRemessaCmd.Flags().StringVarP(&origemID, "origemID", "o", "", "ID da origem (obrigatório)")
	novoGuiaRemessaCmd.Flags().StringVarP(&destinoID, "destinoID", "d", "", "ID do destino (obrigatório)")
	novoGuiaRemessaCmd.Flags().IntVarP(&gprodutoID, "produtoID", "p", 0, "ID do produto (obrigatório)")
	novoGuiaRemessaCmd.Flags().IntVarP(&gquantidade, "quantidade", "q", 0, "Quantidade do produto (obrigatório)")

	// Marcando flags como obrigatórias
	novoGuiaRemessaCmd.MarkFlagRequired("origemID")
	novoGuiaRemessaCmd.MarkFlagRequired("destinoID")
	novoGuiaRemessaCmd.MarkFlagRequired("produtoID")
	novoGuiaRemessaCmd.MarkFlagRequired("quantidade")
}
