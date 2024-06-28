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
		Short: "Internamento de Pacientes",
	}

	internarCmd := &cobra.Command{
		Use:   "internar",
		Short: "Internar paciente",
		Long: `Internar paciente.
			Este comando requer três argumentos, o ID do paciente, o nome e a raça.
			Ele retorna um erro se o paciente não puder ser internado.`,

		Args: cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			pacienteId := args[0]
			pacienteNome := args[1]
			pacienteRaca := args[2]

			repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes.csv")
			servico := application.NovoPaciente(repo)
			err := servico.InternarPaciente(pacienteId, pacienteNome, pacienteRaca)

			if err != nil {
				fmt.Println("Erro ao internar paciente:", err)
				return
			}
			fmt.Println("Paciente internado com sucesso!")
		},
	}

	consultarCmd := &cobra.Command{
		Use:   "consultar",
		Short: "Consultar paciente",
		Long: `Consultar paciente.
			Este comando requer um argumento, o ID do paciente.
			Ele retorna as informações do paciente se encontrado.`,

		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pacienteId := args[0]

			repo := inmem.NovoRepositorioemMemoriaPaciente("pacientes.csv")
			servico := application.NovoPaciente(repo)
			paciente, err := servico.ConsultarPaciente(pacienteId)

			if err != nil {
				fmt.Println("Erro ao consultar paciente:", err)
				return
			}
			fmt.Println("Paciente encontrado:")
			fmt.Println("ID:", paciente.ID)
			fmt.Println("Nome:", paciente.Nome)
			fmt.Println("Raça:", paciente.Raca)
			fmt.Println("Status:", paciente.Status)
		},
	}

	rootCmd.AddCommand(internarCmd)
	rootCmd.AddCommand(consultarCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
