package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-gopher-cli",
		Short: "Gopher CLI in Go",
		Long:  `Gopher CLI application written in Go.`,
	}

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "This command will get the desired Gopher",
		Long:  `This get command will call GitHub repository in order to return the desired Gopher.`,

	

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Por favor insere o ID do paciente e o nome")
				return
			}
			pacienteID := args[0]
			pacienteNome := args[1]

			repo := inmem.NovoRepositorioemMemoriaPaciente()
			servico := application.NovoPaciente(repo)

			err := servico.InternarPaciente(pacienteID, pacienteNome)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Paciente internado com sucesso!")
		},
	}

	rootCmd.AddCommand(getCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}