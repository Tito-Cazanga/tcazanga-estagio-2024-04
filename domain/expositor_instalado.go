package domain

type ExpositorInstalado struct {
	ExpositorID  string
	Localizacao  string
	Quantidade   int
	Produtos     map[int]int
}

func NovoExpositorInstalado(expositorID, localizacao string, produtos map[int]int) *ExpositorInstalado {
	return &ExpositorInstalado{
		ExpositorID: expositorID,
		Localizacao: localizacao,
		Produtos:    produtos,
		Quantidade:  calcularQuantidadeTotal(produtos),
	}
}

func calcularQuantidadeTotal(produtos map[int]int) int {
	total := 0
	for _, quantidade := range produtos {
		total += quantidade
	}
	return total
}
