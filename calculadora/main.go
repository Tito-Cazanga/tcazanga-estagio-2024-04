package calc

type Calculadora string

func New() Calculadora {
	return "0"
}

func (c Calculadora) Ecra() string {
	return string(c)
}


func (c *Calculadora) Dois() {
	c.pressionar("2")
}

func (c *Calculadora) Cinco() {
	c.pressionar("5")
}

func (c *Calculadora) Sete() {
	c.pressionar("7")
}

func (c *Calculadora) Nove() {
	c.pressionar("9")
}

func (c *Calculadora) Mais() {
	c.pressionar("+")
}

func (c *Calculadora) pressionar(s string) {
	if *c == "0" {
		*c = Calculadora(s)
	} else {
		*c = *c + Calculadora(s)
	}
}

//	CODIGO FEITOS PELOS ESTAGIARIOS
func (c *Calculadora) Um() {
	c.pressionar("1")
}


func (c *Calculadora) Menos(){
	c.pressionar("-")
}


func (c *Calculadora) OperacaoMais() {
	c.pressionar("5")
	c.pressionar("7")
}
func (c *Calculadora) OperacaoMenos() {
	c.pressionar("2")
	c.pressionar("9")
}
