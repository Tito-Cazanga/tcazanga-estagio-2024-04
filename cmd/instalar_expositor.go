package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"fitness/domain"
)

var iexpositorID string
var localizacao string
var produtosStr string

var instalarExpositorCmd = &cobra.Command{
	Use:   "instalar-expositor",
	Short: "Instala um novo expositor com produtos iniciais",
	Long:  `Instala um novo expositor em uma localização específica com produtos iniciais.`,
	Args:  cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		produtos := make(map[int]int)

		for _, produto := range strings.Split(produtosStr, ",") {
			parts := strings.Split(produto, ":")
			if len(parts) != 2 {
				fmt.Println("Formato de produto inválido. Use: produtoID:quantidade")
				return
			}
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

		expositorInstalado := domain.NovoExpositorInstalado(iexpositorID, localizacao, produtos)
		fmt.Printf("Expositor instalado:\n")
		fmt.Printf("ExpositorID: %s\n", expositorInstalado.ExpositorID)
		fmt.Printf("Localizacao: %s\n", expositorInstalado.Localizacao)
		fmt.Printf("Produtos: %+v\n", expositorInstalado.Produtos)
	},
}

func init() {
	rootCmd.AddCommand(instalarExpositorCmd)

	// Definindo flags para o comando
	instalarExpositorCmd.Flags().StringVarP(&iexpositorID, "expositorID", "e", "", "ID do expositor (obrigatório)")
	instalarExpositorCmd.Flags().StringVarP(&localizacao, "localizacao", "l", "", "Localização do expositor (obrigatório)")
	instalarExpositorCmd.Flags().StringVarP(&produtosStr, "produtos", "p", "", "Produtos iniciais no formato produtoID:quantidade, separados por vírgula (obrigatório)")

	// Marcando flags como obrigatórias
	instalarExpositorCmd.MarkFlagRequired("expositorID")
	instalarExpositorCmd.MarkFlagRequired("localizacao")
	instalarExpositorCmd.MarkFlagRequired("produtos")
}
