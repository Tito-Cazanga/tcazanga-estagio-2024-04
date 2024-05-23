package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "minha-cli",
	Short: "CLI para gerenciamento de remessas",
	Long:  `Uma CLI para criar e registrar guias de remessa entre expositores.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Podemos adicionar flags globais aqui
}
