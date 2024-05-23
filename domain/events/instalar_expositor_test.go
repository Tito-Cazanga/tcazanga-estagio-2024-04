package instalar_expositor_test

import (
	instalar_expositor "fitness/domain/events"
	"testing"
)

func TestInstalarExpositor(t *testing.T) {
	t.Run("Instalação normal do expositor", func(t *testing.T) {
		ginasioID := 1
		localizacao := "Área de Treino"
		expositor := instalar_expositor.InstalarExpositor(ginasioID, localizacao)
		if expositor == nil {
			t.Errorf("Erro: O expositor não foi instalado corretamente.")
		}
		if expositor.Localizacao != localizacao {
			t.Errorf("Erro: Localização do expositor não corresponde. Esperado: %s, Obtido: %s", localizacao, expositor.Localizacao)
		}
		if expositor.Status != "Instalado" {
			t.Errorf("Erro: O status do expositor não está correto. Esperado: Instalado, Obtido: %s", expositor.Status)
		}
	})

	t.Run("Instalação em outra localização", func(t *testing.T) {
		ginasioID := 1
		localizacao2 := "Recepção"
		expositor2 := instalar_expositor.InstalarExpositor(ginasioID, localizacao2)
		if expositor2.Localizacao != localizacao2 {
			t.Errorf("Erro: Localização do expositor não corresponde. Esperado: %s, Obtido: %s", localizacao2, expositor2.Localizacao)
		}
	})

	t.Run("Instalação em um ginásio diferente", func(t *testing.T) {
		ginasioID2 := 1
		localizacao := "Área de Treino"
		expositor3 := instalar_expositor.InstalarExpositor(ginasioID2, localizacao)
		if expositor3.ID != ginasioID2 {
			t.Errorf("Erro: ID do ginásio do expositor não corresponde. Esperado: %d, Obtido: %d", ginasioID2, expositor3.ID)
		}
	})
}
