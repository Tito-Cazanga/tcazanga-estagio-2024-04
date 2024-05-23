package domain

import "errors"

// Estrutura do Comando de Remessa
type GuiaRemessa struct {
	OrigemID   string
	DestinoID  string
	ProdutoID  int
	Quantidade int
}

// Estrutura do Evento de Remessa
type RemessaRegistrada struct {
	OrigemID   string
	DestinoID  string
	ProdutoID  int
	Quantidade int
}

// Novo Guia de Remessa
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

// Registrar Remessa
func (cmd *GuiaRemessa) RegistrarRemessa(origem, destino *Expositor) (*RemessaRegistrada, error) {
	quantidadeEstoque, existe := origem.Estoque[cmd.ProdutoID]
	if !existe {
		return nil, errors.New("produto não encontrado no estoque da origem")
	}

	if quantidadeEstoque < cmd.Quantidade {
		return nil, errors.New("estoque insuficiente na origem")
	}

	// Deduzir do estoque da origem
	origem.Estoque[cmd.ProdutoID] -= cmd.Quantidade

	// Adicionar ao estoque do destino
	if destino.Estoque == nil {
		destino.Estoque = make(map[int]int)
	}
	destino.Estoque[cmd.ProdutoID] += cmd.Quantidade

	return &RemessaRegistrada{
		OrigemID:   cmd.OrigemID,
		DestinoID:  cmd.DestinoID,
		ProdutoID:  cmd.ProdutoID,
		Quantidade: cmd.Quantidade,
	}, nil
}
