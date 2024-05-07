package inventario_test

import (
	"testing"
)

func TestAdicionarProduto(t *testing.T) {
    // Testa a função AddProduct para garantir que um produto seja adicionado corretamente ao inventário
    // Setup
    inventario := NewInventario()
    produto := Produto{Nome: "Produto X", Categoria: "Categoria A", Validade: "2024-05-31"}
    
    // Execução
    inventario.AdicionarProduto(produto)
    
    // Verificação
    if len(inventario.Produtos) != 1 {
        t.Errorf("Esperava-se que o inventário tivesse 1 produto, mas encontrou %d", len(inventario.Produtos))
    }
}

// Testes similares para outras funcionalidades de gerenciamento de inventário, como UpdateProduct e RemoveProduct
