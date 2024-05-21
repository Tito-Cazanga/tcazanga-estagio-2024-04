package application

import (
	"errors"
	"fitness/domain"
)

type EmitirFatura struct {
	GinasioID string
	Consumos  []*domain.ConsumoRegistrado
}

var ErrListaConsumosVazia = errors.New("a lista de consumos está vazia")

func (comando *EmitirFatura) NewFatura() (*domain.Fatura, error) {
	// Verificar se há consumos
	if len(comando.Consumos) == 0 {
		return nil, ErrListaConsumosVazia
	}

	// Gerar a fatura usando a função da camada de domínio
	return domain.FaturaEmitida(comando.GinasioID, comando.Consumos)
}
