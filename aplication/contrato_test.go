package application

import (
	"testing"
)

func TestContract(t *testing.T){
	t.Run("Cria contrato", func(t *testing.T){
		// Arrange
		expected := "Contrato foi criado"

		// Act
		result := NewContrato()
		
		// Assert
		if expected != string(result) {
			t.Error("Contrato não exite")
		}
	})
}