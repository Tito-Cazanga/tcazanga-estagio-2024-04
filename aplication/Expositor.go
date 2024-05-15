package application

type Expositor struct {
	ID 				string
	Localizacao 	string
	Estoque 		map[int]int

}

type ExpositorAbastecido struct {
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}

type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID 	int
	Quantidade 	int
}

type ConsumoRegistrado struct {
	GinásioID  string
	ProdutoID  int
	Quantidade int
}

type ConsumirProduto struct {
	ExpositorID int
	ProdutoID   int
	Quantidade  int
}

func (comando *AbastecerExpositor) ExecutarAbastecerExpositor(expositor *Expositor) *ExpositorAbastecido {
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

func (cmd *ConsumirProduto) Executar(expositor *Expositor) *ConsumoRegistrado {
	// Verificar se o expositor possui o produto em estoque
	quantidadeEstoque, ok := expositor.Estoque[cmd.ProdutoID]
	if !ok {
		return nil
	}

	if quantidadeEstoque < cmd.Quantidade {
		return nil
	}

	expositor.Estoque[cmd.ProdutoID] -= cmd.Quantidade

	return &ConsumoRegistrado{
		GinásioID:  expositor.ID,
		ProdutoID:  cmd.ProdutoID,
		Quantidade: cmd.Quantidade,
	}
}
