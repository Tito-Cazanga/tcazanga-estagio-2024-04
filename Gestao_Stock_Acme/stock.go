package stock

type Produto struct{
	ID int
	Categoria string
	Detalhe string
	EstadoProduto string
	Validade string
}
func (p Produto) Apresentacao() string{
	return ""
}
func(p *Produto) Stock(s Produto){
	p.AdicionarProduto(Produto{ID: 13,
		Categoria: "AA",
		Detalhe: "bb",
		EstadoProduto: "BB",
		Validade: "CC",
		},
	)
}
func (p *Produto)AdicionarProduto(d Produto){
	
}


