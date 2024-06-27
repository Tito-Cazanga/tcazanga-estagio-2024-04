package main

import (
	"fmt"
	"os"

	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/adapter/inmem"
	"github.com/Tito-Cazanga/tcazanga-estagio-2024-04/application"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cvl",
		Short: "Aplicação CLI para internar paciente.",
		Long:  `Aplicação CLI para consulta paciente usando Cobra em Go.`,
	}

	
	consultarCmd := &cobra.Command{
		Use:   "consulta_paciente",
		Short: "Este comando irá consultar paciente",
		Long:  `Este comando consulta_paciente verifica se o paciente esta internado e retorna se sim ou nao.`,

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
			fmt.Println("Paciente esta internado internado!")
		},
	}

	rootCmd.AddCommand(consultarCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
