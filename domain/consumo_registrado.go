package domain

import "errors"

type ConsumoRegistrado struct {
	GinásioID  string
	ProdutoID  int
	Quantidade int
}

func NewConsumoRegistrado(ginásioID string, produtoID, quantidade int) (*ConsumoRegistrado, error) {
	if quantidade <= 0 {
		return nil, errors.New("a quantidade deve ser positiva")
	}

	return &ConsumoRegistrado{
		GinásioID:  ginásioID,
		ProdutoID:  produtoID,
		Quantidade: quantidade,
	}, nil
}