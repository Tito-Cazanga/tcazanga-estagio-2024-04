package main

import (
	"fmt"
	"sort"
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

func (l *Lote) DiasParaValidade() int {
    hoje := time.Now().Truncate(24 * time.Hour)
    diferenca := l.DataDeValidade.Sub(hoje)
    return int(diferenca.Hours() / 24)
}

func RecomendarLotes(lotes []*Lote) []*Lote {
    
    // Copia os lotes para não modificar o original
    lotesCopia := make([]*Lote, len(lotes))
    copy(lotesCopia, lotes)

    // Ordena os lotes com base na proximidade da validade
    sort.Slice(lotesCopia, func(i, j int) bool {
        return lotesCopia[i].DiasParaValidade() < lotesCopia[j].DiasParaValidade()
    })

    return lotesCopia
}

func main() {
    lote1 := NovoLote(1, "Lote001", 100, "CorredorA", "Armario1", "Prateleira2", 2024, 6, 1)
    lote2 := NovoLote(2, "Lote002", 50, "CorredorB", "Armario2", "Prateleira1", 2024, 6, 15)
    lote3 := NovoLote(3, "Lote003", 200, "CorredorC", "Armario3", "Prateleira3", 2026, 5, 20)

    lotes := []*Lote{lote1, lote2, lote3}

    // Recomendação dos lotes com data de validade mais próxima
    lotesRecomendados := RecomendarLotes(lotes)

    fmt.Println("Lotes recomendados com data de validade mais próxima:")
    for _, lote := range lotesRecomendados {
        fmt.Printf("Lote: %s, Dias para validade: %d\n", lote.Lote, lote.DiasParaValidade())
    }
}
