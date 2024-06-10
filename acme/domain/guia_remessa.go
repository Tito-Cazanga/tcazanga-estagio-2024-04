package domain

import (
	"errors"
	"time"
)

type GuiaRemessa struct {
	OrigemID   string
	DestinoID  string
	ProdutoID  int
	Quantidade int
}


func NovoGuiaRemessa(origemID, destinoID string, produtoID, quantidade int) (*GuiaRemessa, error) {
	if quantidade <= 0 {
		return nil, errors.New("a quantidade deve ser positiva")
	}
	return &GuiaRemessa{
		OrigemID:   origemID,
		DestinoID:  destinoID,
		ProdutoID:  produtoID,
		Quantidade: quantidade,
	}, nil
}

type RemessaCriada struct {
	OrigemID   string
	DestinoID  string
	ProdutoID  int
	Quantidade int
	Data       time.Time
}

var ErrLimiteAbastecimentosAtingido = errors.New("limite de abastecimentos semanais atingido")

func (cmd *GuiaRemessa) CriarRemessa(origem, destino *Expositor) (*RemessaCriada, error) {

	quantidadeEstoqueOrigem := origem.Estoque[cmd.ProdutoID]

	if quantidadeEstoqueOrigem < cmd.Quantidade {
		return nil, errors.New("estoque insuficiente na origem")
	}

	if origem.ID == destino.ID {
		return nil, errors.New("a origem e o destino não podem ser o mesmo expositor")
	}

	if err := verificarLimiteAbastecimentos(destino); err != nil {
		return nil, err
	}


	remessa := &RemessaCriada{
		OrigemID:   origem.ID,
		DestinoID:  destino.ID,
		ProdutoID:  cmd.ProdutoID,
		Quantidade: cmd.Quantidade,
		Data:       time.Now(),
	}

	return remessa, nil
}

func verificarLimiteAbastecimentos(expositor *Expositor) error {
	semanaAtual := time.Now().AddDate(0, 0, -int(time.Now().Weekday()))

	var abastecimentosRecentes []time.Time
	for _, data := range expositor.Abastecimentos {
		if data.After(semanaAtual) {
			abastecimentosRecentes = append(abastecimentosRecentes, data)
		}
	}

	if len(abastecimentosRecentes) >= 2 {
		return ErrLimiteAbastecimentosAtingido
	}

	expositor.Abastecimentos = append(abastecimentosRecentes, time.Now())
	return nil
}