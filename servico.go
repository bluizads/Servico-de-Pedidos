package main

import "fmt"

type ServicoPedido struct {
	repoProduto RepositorioProduto
	repoPedido  RepositorioPedido
	contadorID  int
}

func NovoServicoPedido(rProd RepositorioProduto, rPed RepositorioPedido) *ServicoPedido {
	return &ServicoPedido{
		repoProduto: rProd,
		repoPedido:  rPed,
	}
}

func (serv *ServicoPedido) CriarPedido(cliente string, itens []ItemPedido) (*Pedido, error) {
	// validar
	if cliente == "" {
		return nil, ErrClienteInvalido
	}
	if len(itens) == 0 {
		return nil, ErrPedidoVazio
	}
	for _, item := range itens {
		if item.Quantidade <= 0 {
			return nil, ErrQuantidadeInvalida
		}
	}

	// tem estoque?
	for _, item := range itens {
		produto, err := serv.repoProduto.BuscarPorID(item.ProdutoId)
		if err != nil {
			return nil, err
		}
		if produto.Estoque < item.Quantidade {
			return nil, ErrEstoqueInsuficiente
		}
	}

	itensDoPedido := make([]ItemPedido, 0, len(itens))

	// reduz estoque
	for _, item := range itens {
		produto, err := serv.repoProduto.BuscarPorID(item.ProdutoId)
		if err != nil {
			return nil, err
		}

		err = produto.ReduzirEstoque(item.Quantidade)
		if err != nil {
			return nil, err
		}

		itemFinal := ItemPedido{
			ProdutoId:     item.ProdutoId,
			PrecoNaCompra: produto.Preco,
			Quantidade:    item.Quantidade,
		}

		itensDoPedido = append(itensDoPedido, itemFinal)
	}

	serv.contadorID++
	id := fmt.Sprintf("PED-%03d", serv.contadorID)

	pedido := Pedido{
		ID:      id,
		Cliente: cliente,
		Itens:   itensDoPedido,
		Status:  StatusPendente,
	}

	serv.repoPedido.Salvar(pedido)

	return &pedido, nil
}
