package domain

type Expositor struct {
	ID          string
	Localizacao string
	Estoque     map[int]int
}
type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

type InstalarExpositor struct {
	ExpositorID  string
	Localizacao  string
	Quantidade   int
	Produtos     map[int]int
}


func (cmd *AbastecerExpositor) AbastecerExpositor(expositor *Expositor) *AbastecerExpositor {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}

	expositor.Estoque[cmd.ProdutoID] += cmd.Quantidade

	return &AbastecerExpositor{
		ExpositorID: cmd.ExpositorID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}
}

func NovoExpositorInstalado(expositorID, localizacao string, produtos map[int]int) *InstalarExpositor {
	return &InstalarExpositor{
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


