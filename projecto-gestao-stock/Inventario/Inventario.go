package inventario

import (
    "time"

)

type Produto struct {
	ID             int
	Nome           string
}

func NovoProduto(id int, nome string) *Produto {
	return &Produto{
		ID:             id,
		Nome:           nome,
	}
}

type Lote struct {
    ProdutoID       int
    Lote            string
    Quantidade      int
    Localizacao     Localizacao
    DataDeValidade  time.Time
}

type Localizacao struct {
    Corredor    string
    Armario     string
    Prateleira  string
}

func NovoLote(produtoID int, lote string, quantidade int, corredor, armario, prateleira string, ano, mes, dia int) *Lote {
    dataDeValidade := time.Date(ano, time.Month(mes), dia, 0, 0, 0, 0, time.UTC)

    localizacao := Localizacao{
        Corredor:    corredor,
        Armario:     armario,
        Prateleira:  prateleira,
    }

    return &Lote{
        ProdutoID:      produtoID,
        Lote:           lote,
        Quantidade:     quantidade,
        Localizacao:    localizacao,
        DataDeValidade: dataDeValidade,
    }
}


