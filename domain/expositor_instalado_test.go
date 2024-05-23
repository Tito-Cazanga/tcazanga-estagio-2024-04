package domain_test

import (
	"fitness/domain"
	"reflect"
	"testing"
)

func TestNovoExpositorInstalado(t *testing.T) {
	expositorID := "1"
	localizacao := "Sala de Musculação"
	produtos := map[int]int{1: 50, 2: 30}

	expectedEvento := &domain.ExpositorInstalado{
		ExpositorID: expositorID,
		Localizacao: localizacao,
		Produtos:    produtos,
		Quantidade:  80,
	}

	evento := domain.NovoExpositorInstalado(expositorID, localizacao, produtos)

	if !reflect.DeepEqual(evento, expectedEvento) {
		t.Errorf("Evento esperado: %+v, mas obtido: %+v", expectedEvento, evento)
	}
}
