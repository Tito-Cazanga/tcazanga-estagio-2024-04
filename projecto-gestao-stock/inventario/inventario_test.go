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
