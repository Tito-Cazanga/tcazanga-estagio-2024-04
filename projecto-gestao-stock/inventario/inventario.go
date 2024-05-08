// inventario.go
package inventario

import (
	"fmt"
	"sort"
	"time"
)

type Produto struct {
	Nome        string
	DataValidade time.Time
}

type Inventorio []Produto

func (i Inventorio) OrganizarPorDataValidade() {
	sort.Slice(i, func(j, k int) bool {
		return i[j].DataValidade.Before(i[k].DataValidade)
	})
}

func (i Inventorio) RecomendarProdutosProximosAVencerimento() []Produto {
	var proximosAVencerimento []Produto
	for _, produto := range i {
		if produto.DataValidade.Sub(time.Now()) < 30*24*time.Hour {
			proximosAVencerimento = append(proximosAVencerimento, produto)
		}
	}
	return proximosAVencerimento
}

