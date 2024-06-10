package domain

import (
	"errors"
	"time"
)

type Expositor struct {
	ID             string
	Localizacao    string
	Estoque        map[int]int
	Abastecimentos []time.Time
}

type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

type InstalarExpositor struct {
	ExpositorID string
	Localizacao string
	Produtos    map[int]int
}

type AbastecimentoRealizado struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

type ExpositorInstalado struct {
	ExpositorID string
	Localizacao string
	Produtos    map[int]int
}

func (cmd *AbastecerExpositor) Abastecer(expositor *Expositor) *AbastecimentoRealizado {
	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}

	expositor.Estoque[cmd.ProdutoID] += cmd.Quantidade

	return &AbastecimentoRealizado{
		ExpositorID: cmd.ExpositorID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}
}

func (cmd *AbastecerExpositor) AbastecerExpositor(expositor *Expositor) (*AbastecerExpositor, error) {
	if len(expositor.Abastecimentos) >= 2 && numAbastecimentosEstaSemana(expositor.Abastecimentos) >= 2 {
		return nil, errors.New("o expositor já foi abastecido duas vezes esta semana")
	}

	if expositor.Estoque == nil {
		expositor.Estoque = make(map[int]int)
	}

	expositor.Estoque[cmd.ProdutoID] += cmd.Quantidade
	expositor.Abastecimentos = append(expositor.Abastecimentos, time.Now())

	return &AbastecerExpositor{
		ExpositorID: cmd.ExpositorID,
		ProdutoID:   cmd.ProdutoID,
		Quantidade:  cmd.Quantidade,
	}, nil
}

func numAbastecimentosEstaSemana(abastecimentos []time.Time) int {
	count := 0
	now := time.Now()
	_, currentWeek := now.ISOWeek()

	for _, data := range abastecimentos {
		_, week := data.ISOWeek()
		if data.Year() == now.Year() && week == currentWeek {
			count++
		}
	}

	return count
}

func NovoExpositorInstalado(expositorID, localizacao string, produtos map[int]int) *ExpositorInstalado {
	return &ExpositorInstalado{
		ExpositorID: expositorID,
		Localizacao: localizacao,
		Produtos:    produtos,
	}
}