package inventario

import (
    "time"

)

type Produto struct {
	ID             int
	Nome           string
	DataDeValidade time.Time
}

func NovoProduto(id int, nome string, dataDeValidade time.Time) *Produto {
	return &Produto{
		ID:             id,
		Nome:           nome,
		DataDeValidade: dataDeValidade,
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
    dataDeValidade := time.Now().AddDate(ano, mes, dia)
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


