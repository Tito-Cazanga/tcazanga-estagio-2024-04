package domain_test

import (
	"fitness/domain"
	"testing"
)

func TestNovoExpositorInstalado(t *testing.T) {
	expositorID := "1"
	localizacao := "Sala de Musculação"
	produtos := map[int]int{1: 50, 2: 30}

	expectedEvento := &domain.InstalarExpositor{
		ExpositorID: expositorID,
		Localizacao: localizacao,
		Produtos:    produtos,
		Quantidade:  80,
	}
	

	evento := domain.NovoExpositorInstalado(expositorID, localizacao, produtos)

	if evento == expectedEvento {
		t.Errorf("Evento esperado: %+v, mas obtido: %+v", expectedEvento, evento)
	}
}
