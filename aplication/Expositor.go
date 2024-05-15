package application

type Produto struct{
	ID int

}

type Expositor struct {
	ID 				string
	Localizacao 	string
	Estoque 		map[int]int

}

type ExpositorAbastecido struct{
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}

type AbastecerExpositor struct{
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}


func (comando *AbastecerExpositor) Executar(expositor *Expositor) *ExpositorAbastecido {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}
	
	expositor.Estoque[comando.ProdutoID] += comando.Quantidade

	return &ExpositorAbastecido{
		ExpositorID: comando.ExpositorID,
		ProdutoID: comando.ProdutoID,
		Quantidade: comando.Quantidade,
	}
}