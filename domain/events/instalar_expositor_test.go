package instalar_expositor_test

import (
	instalar_expositor "fitness/domain/events"
	"testing"
)

// Teste para verificar se o expositor é instalado corretamente
func TestInstalarExpositor(t *testing.T) {
	// Simulando a instalação do expositor em um ginásio
	ginasioID := 1
	localizacao := "Área de Treino"
	expositor := instalar_expositor.InstalarExpositor(ginasioID, localizacao)

	// Verificando se o expositor foi instalado corretamente
	if expositor == nil {
		t.Errorf("Erro: O expositor não foi instalado corretamente.")
	}

	if expositor.Localizacao != localizacao {
		t.Errorf("Erro: Localização do expositor não corresponde. Esperado: %s, Obtido: %s", localizacao, expositor.Localizacao)
	}

	if expositor.Status != "Instalado" {
		t.Errorf("Erro: O status do expositor não está correto. Esperado: Instalado, Obtido: %s", expositor.Status)
	}
}
