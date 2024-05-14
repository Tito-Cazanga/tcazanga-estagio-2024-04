package application_test

import (
	application "fitness/aplication"
	"testing"

	
)

func verificar(t *testing.T, esperadoExpositor, resultExpositor string) {
	if esperadoExpositor != resultExpositor {
		t.Errorf("Recebido %s, esperado %s ", resultExpositor, esperadoExpositor)
	}
}

func TestAbastecerExpositor(t *testing.T) {

	t.Run("Abastecer Expositor", func(t *testing.T) {
		//Arrange
		esperadoExpositor := "Expositor abastecido"

		//Act
		resultExpositor := application.NewExpositor()

		//Assert
		verificar(t, esperadoExpositor, string(resultExpositor))

	})
}

func TestExisteProduto(t *testing.T) {
	t.Run("Existencia de Produto", func(t *testing.T) {

		//arrange
		produto := []application.Produto{
			{ID: "2"},
			{ID: "1"},
		}

		//act
		//assert
		if len(produto) != 2 {
			t.Error("Produto Não Existe")
		}

	})

	t.Run("Existencia de um produto especifico", func(t *testing.T) {

		//Arrange
		produto := application.Produto{ID: "2",}
		
		//Act
		existe := application.ExisteProduto(produto, "2")

		//Assert
		if !existe {
			t.Error("O produto requerido não existe")
		}
	})

}

