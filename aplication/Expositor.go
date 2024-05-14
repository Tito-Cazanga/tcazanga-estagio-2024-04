package application

type AbasterExpositor string

func NewExpositor() AbasterExpositor {
	return "Expositor abastecido"
}

type Produto struct{
	ID string
}

func ExisteProduto(produto Produto, produtoID string) bool {
	
	if produto.ID == produtoID {
		return true
	}
	
	return false
}