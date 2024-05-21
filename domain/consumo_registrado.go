package domain

import "errors"

type ConsumoRegistrado struct {
	GinasioID  string
	ProdutoID  int
	Quantidade int
}

func NewConsumoRegistrado(ginasioID string, produtoID, quantidade int) (*ConsumoRegistrado, error) {
	if quantidade <= 0 {
		return nil, errors.New("a quantidade deve ser positiva")
	}

	return &ConsumoRegistrado{
		GinasioID:  ginasioID,
		ProdutoID:  produtoID,
		Quantidade: quantidade,
	}, nil
}