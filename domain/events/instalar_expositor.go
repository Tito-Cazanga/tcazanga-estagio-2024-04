package instalar_expositor

// Definindo a estrutura para o expositor
type Expositor struct {
	ID         int
	Localizacao string
	Status     string
	// Outros campos necessários podem ser adicionados aqui
}

// Função para instalar um novo expositor em um ginásio
func InstalarExpositor(ginasioID int, localizacao string) *Expositor {
	// Lógica para instalação do expositor...
	expositor := &Expositor{
		ID:         1, // ID do expositor
		Localizacao: localizacao,
		Status:     "Instalado",
	}
	return expositor
}
