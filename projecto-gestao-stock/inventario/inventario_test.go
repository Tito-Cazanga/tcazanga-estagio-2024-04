// inventario_test.go
package inventario

import (
	"testing"
	"time"
)

func TestOrganizarPorDataValidade(t *testing.T) {
	inventario := Inventorio{
		{Nome: "Produto 1", DataValidade: time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 2", DataValidade: time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 3", DataValidade: time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC)},
	}

	inventario.OrganizarPorDataValidade()

	expected := Inventorio{
		{Nome: "Produto 2", DataValidade: time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 1", DataValidade: time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 3", DataValidade: time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC)},
	}

	if!equalInventorio(inventario, expected) {
		t.Errorf("Inventorio não foi organizado corretamente")
	}
}

func TestRecomendarProdutosProximosAVencerimento(t *testing.T) {
	inventario := Inventorio{
		{Nome: "Produto 1", DataValidade: time.Date(2023, 3, 15, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 2", DataValidade: time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC)},
		{Nome: "Produto 3", DataValidade: time.Date(2023, 4, 20, 0, 0, 0, 0, time.UTC)},
	}

	proximosAVencerimento := inventario.RecomendarProdutosProximosAVencerimento()

	expected := []Produto{
		{Nome: "Produto 2", DataValidade: time.Date(2023, 2, 10, 0, 0, 0, 0, time.UTC)},
	}

	if!equalProdutos(proximosAVencerimento, expected) {
		t.Errorf("Produtos proximos a vencerimento não foram recomendados corretamente")
	}
}

func equalInventorio(a, b Inventorio) bool {
	if len(a)!= len(b) {
		return false
	}
	for i := range a {
		if a[i].Nome!= b[i].Nome || a[i].DataValidade!= b[i].DataValidade {
			return false
		}
	}
	return true
}

func equalProdutos(a, b []Produto) bool {
	if len(a)!= len(b) {
		return false
	}
	for i := range a {
		if a[i].Nome!= b[i].Nome || a[i].DataValidade!= b[i].DataValidade {
			return false
		}
	}
	return true
}