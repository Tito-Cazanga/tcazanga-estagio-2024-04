package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fitness-cli",
	Short: "CLI para gestão de actividades na empresa Fitness",
	Long:  `CLI para gestão de todas actividades na empresa Fitness apartir da linha de comando.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

