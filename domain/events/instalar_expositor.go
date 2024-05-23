package instalar_expositor


type Expositor struct {
	ID         int
	Localizacao string
	Status     string

}

func InstalarExpositor(ginasioID int, localizacao string) *Expositor {
	expositor := &Expositor{
		ID:         1,
		Localizacao: localizacao,
		Status:     "Instalado",
	}
	return expositor
}
