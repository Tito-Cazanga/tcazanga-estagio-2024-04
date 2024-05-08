package inventario

import "time"

// Produto representa um produto com um ID, nome e data de validade.
type Produto struct {
	ID             int
	Nome           string
	DataDeValidade time.Time
}

// NovoProduto cria e retorna um novo produto com os valores fornecidos.
func NovoProduto(id int, nome string, dataDeValidade time.Time) *Produto {
	return &Produto{
		ID:             id,
		Nome:           nome,
		DataDeValidade: dataDeValidade,
	}
}
