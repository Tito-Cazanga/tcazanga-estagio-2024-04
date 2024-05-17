
package domain

import (
	"errors"
	"fmt"
	"time"
)

type Fatura struct {
	ID           string
	DataEmissao  time.Time
	GinásioID    string
	Total        float64
	Consumos     []*ConsumoRegistrado
}

func FaturaEmitida(ginásioID string, consumos []*ConsumoRegistrado) (*Fatura, error) {
	if len(consumos) == 0 {
		return nil, errors.New("a lista de consumos não pode estar vazia")
	}

	total := 0.0
	for _, consumo := range consumos {
		precoProduto := 1.0
		total += float64(consumo.Quantidade) * precoProduto
	}

	return &Fatura{
		ID:           generateID(),     
		DataEmissao:  time.Now(),      
		GinásioID:    ginásioID,
		Total:        total,
		Consumos:     consumos,
	}, nil
}

func generateID() string {
	// Implementação fictícia para gerar um ID único
	return fmt.Sprintf("%d", time.Now().UnixNano())
}