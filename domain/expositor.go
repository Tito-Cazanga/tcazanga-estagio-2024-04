package domain

import (
	"errors"
	"time"
)

// Expositor representa um expositor com uma localização e um estoque de produtos.
type Expositor struct {
	ID          string
	Localizacao string
	Estoque     map[int]int
	Abastecimentos  []time.Time
}

// AbastecerExpositor é o comando para abastecer um expositor com um determinado produto.
type AbastecerExpositor struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

// InstalarExpositor é o comando para instalar um novo expositor com produtos iniciais.
type InstalarExpositor struct {
	ExpositorID string
	Localizacao string
	Produtos    map[int]int
}

// AbastecimentoRealizado é o evento gerado após um expositor ser abastecido.
type AbastecimentoRealizado struct {
	ExpositorID string
	ProdutoID   int
	Quantidade  int
}

// ExpositorInstalado é o evento gerado após um novo expositor ser instalado.
type ExpositorInstalado struct {
	ExpositorID string
	Localizacao string
	Produtos    map[int]int
}

// Abastecer executa o comando de abastecimento no expositor e retorna o evento gerado.
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




// Abastecer Expositor com Restrição Semanal
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

// Função Auxiliar para Verificar Abastecimentos na Semana Corrente
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


// NovoExpositorInstalado cria um novo comando de instalação de expositor e retorna o evento gerado.
func NovoExpositorInstalado(expositorID, localizacao string, produtos map[int]int) *ExpositorInstalado {
	return &ExpositorInstalado{
		ExpositorID: expositorID,
		Localizacao: localizacao,
		Produtos:    produtos,
	}
}

// calcularQuantidadeTotal calcula a quantidade total de produtos.
func calcularQuantidadeTotal(produtos map[int]int) int {
	total := 0
	for _, quantidade := range produtos {
		total += quantidade
	}
	return total
}
