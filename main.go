package main

import (
	"fmt"
)

func main() {
	repoProduto := NovoRepositorioProdutoMemoria()
	repoPedido := NovoRepositorioPedidoMemoria()

	servico := NovoServicoPedido(repoProduto, repoPedido)

	repoProduto.Salvar(Produto{ID: "P001", Nome: "Notebook", Preco: 3500, Estoque: 5})

	itens := []ItemPedido{
		{ProdutoId: "P001", Quantidade: 1},
	}

	pedido, err := servico.CriarPedido("Ana", itens)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("pedido criado:", pedido.ID, "total:", pedido.CalcularTotal())

}
