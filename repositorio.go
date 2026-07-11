package main

// pra ser um repositorio tem q ter
type RepositorioProduto interface {
	Salvar(produto Produto)
	BuscarPorId(id string) (*Produto, error)
	Listar() []Produto
}

//em memória
type RepositorioProdutoMemoria struct {
	produtos map[string]*Produto
}

func NovoRepositorioProdutoMemoria() *RepositorioProdutoMemoria {
	return &RepositorioProdutoMemoria{
		produtos: make(map[string]*Produto),
	}
}

func (repo *RepositorioProdutoMemoria) BuscarPorId(id string) (*Produto, error) {
	produto, ok := repo.produtos[id]
	if !ok {
		return nil, ErrProdutoNaoEncontrado
	}
	return produto, nil
}

// para pedidos
type RepositorioPedido interface {
	Salvar(pedido Pedido)
	BuscarPorId(id string) (*Pedido, error)
	Listar() []Pedido
}

//em memória
type RepositorioPedidoMemoria struct {
	pedidos map[string]*Pedido
}

func NovoRepositorioPedidoMemoria() *RepositorioPedidoMemoria {
	return &RepositorioPedidoMemoria{
		pedidos: make(map[string]*Pedido),
	}
}

func (repo *RepositorioPedidoMemoria) BuscarPorId(id string) (*Pedido, error) {
	pedido, ok := repo.pedidos[id]
	if !ok {
		return nil, ErrPedidoNaoEcncontrado
	}
	return pedido, nil
}

func (r *RepositorioPedidoMemoria) Salvar(pedido Pedido) {
	r.pedidos[pedido.ID] = &pedido
}
func (r *RepositorioPedidoMemoria) Listar() []Pedido {
	lista := make([]Pedido, 0, len(r.pedidos))
	for _, pedido := range r.pedidos {
		lista = append(lista, *pedido)
	}
	return lista
}
