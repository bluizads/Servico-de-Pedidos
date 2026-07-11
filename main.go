package main

import (
	"fmt"
)

func main() {
	repoProduto := NovoRepositorioProdutoMemoria()
	repoPedido := NovoRepositorioPedidoMemoria()

	servico := NovoServicoPedido(repoProduto, repoPedido)

	repoProduto.Salvar(Produto{ID: "P001", Nome: "Notebook", Preco: 3500.00, Estoque: 5})
	repoProduto.Salvar(Produto{ID: "P002", Nome: "Mouse", Preco: 80.00, Estoque: 10})
	repoProduto.Salvar(Produto{ID: "P003", Nome: "Teclado", Preco: 180.00, Estoque: 8})

	itens := []ItemPedido{
		{ProdutoId: "P001", Quantidade: 1},
	}

	pedido, err := servico.CriarPedido("Ana", itens)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("\n>> Pedido criado:", pedido.ID, "total: R$ ", pedido.CalcularTotal())
	fmt.Println("estoque do notebook agora:", repoProduto.produtos["P001"].Estoque)

}
