package calc_test

import (
	calc "calculadora"
	"testing"
)

func verifica(t *testing.T, esperado, actual string) {
	if esperado != actual {
		t.Logf("%s != %s", esperado, actual)
		t.Fail()
	}
}

func TestCria_0(t *testing.T) {
	// Arrange
	// Act
	c := calc.New()

	// Assert
	verifica(t, "0", c.Ecra())
}

func TestPressiona5DepoisDeLigar_5(t *testing.T) {
	// Arrange
	c := calc.New()

	// Act
	c.Cinco()

	// Assert
	verifica(t, "5", c.Ecra())

}

func TestPressionaMaisDepoisDe7_7Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Sete()

	// Act
	c.Mais()

	//Assert
	verifica(t, "7+", c.Ecra())

}

func TestPressionaMaisDepoisDe2_9_7__297Mais(t *testing.T) {
	// Arrange
	c := calc.New()
	c.Dois()
	c.Nove()
	c.Sete()

	// Act
	c.Mais()

	// Assert
	verifica(t, "297+", c.Ecra())
}

//	CODIGO FEITOS PELOS ESTAGIARIOS

func TestPresionarMenosDepoisDe1_1Menos(t *testing.T){
	//Arrange
	 c := calc.New()
	 c.Um()
	 
	 //Act
	 c.Menos()

	 //Assert
	 verifica(t, "1-", c.Ecra())
}

func TestPresionarMenosDepoisDe5_7_2_572Menos(t *testing.T){
	//Arrange
	c := calc.New()
	c.Cinco()
	c.Sete()
	c.Dois()

	//Act
	c.Menos()

	//Assert
	verifica(t, "572-", c.Ecra())
}

func TesteSomaDe7_5_Igual7Mais5(t *testing.T){
	//Arrange
	c := calc.New()
	c.OperacaoMais()
	
	//Act
	c.Mais()
	
	//Assert
	verifica(t, "7+5", c.Ecra())
	
}	

func TesteMenosDe2_9_Igual2Menos9(t *testing.T){
	//Arrange
	c := calc.New()
	c.OperacaoMenos()
	
	//Act
	c.Menos()
	
	//Assert
	verifica(t, "2-9", c.Ecra())
	
}	
