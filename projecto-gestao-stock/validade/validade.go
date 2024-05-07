package validade

import (
	"time"
)

type Produto struct {
	ID                       string
	Descricao                string
	Categoria                string
	AmbienteDeArmazenamento  string
}

type Lote struct {
	ID              string
	ProdutoId       string
	Lote            string
	Validade        time.Time
	Quantidade      int
}

type Inventario struct {
	Produtos []*Produto
	Lotes    []*Lote
}

func NewInventario() *Inventario {
	return &Inventario{}
}

func (iv *Inventario) CriaNovoProduto(produto *Produto) {
	iv.Produtos = append(iv.Produtos, produto)
}

func (iv *Inventario) CriaNovoLote(lote *Lote) {
	iv.Lotes = append(iv.Lotes, lote)
}