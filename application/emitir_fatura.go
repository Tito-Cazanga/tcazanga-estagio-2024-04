package application

import (
	"errors"
	"fitness/domain"
	"fmt"
	"time"
)

type EmitirFatura struct {
	GinasioID string
	Consumos  []*domain.ConsumirProduto
}

type Fatura struct {
	ID          string
	DataEmissao time.Time
	GinasioID   string
	Total       float64
	Consumos    []*domain.ConsumirProduto
}


var ErrListaConsumosVazia = errors.New("a lista de consumos está vazia")

func (comando *EmitirFatura) NewFatura() (*Fatura, error) {
	
	if len(comando.Consumos) == 0 {
		return nil, ErrListaConsumosVazia
	}

	return FaturaEmitida(comando.GinasioID, comando.Consumos)
}

func FaturaEmitida(ginásioID string, consumos []*domain.ConsumirProduto) (*Fatura, error) {
	if len(consumos) == 0 {
		return nil, errors.New("a lista de consumos não pode estar vazia")
	}

	total := 0.0
	for _, consumo := range consumos {
		precoProduto := 1.0
		total += float64(consumo.Quantidade) * precoProduto
	}

	return &Fatura{
		ID:          generateID(),
		DataEmissao: time.Now(),
		GinasioID:   ginásioID,
		Total:       total,
		Consumos:    consumos,
	}, nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
