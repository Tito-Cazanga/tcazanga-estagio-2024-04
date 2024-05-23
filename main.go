package main

import (
	"fitness/cmd"
	"fitness/domain"
	"fmt"
)

func main() {
	cmd.Execute()


	expositorID := "1"
	localizacao := "Sala de Musculação"
	produtos := map[int]int{1: 50, 2: 30} // Exemplo de produtos no expositor

	evento := domain.NovoExpositorInstalado(expositorID, localizacao, produtos)

	fmt.Printf("Evento: %+v\n", evento)


}
