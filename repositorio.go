package main

// pra ser um repositorio tem q ter
type RepositorioProduto interface {
	Salvar(produto Produto)
	BuscarPorID(ID string) (*Produto, error)
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

func (repo *RepositorioProdutoMemoria) BuscarPorID(ID string) (*Produto, error) {
	produto, ok := repo.produtos[ID]
	if !ok {
		return nil, ErrProdutoNaoEncontrado
	}
	return produto, nil
}

func (repo *RepositorioProdutoMemoria) Salvar(produto Produto) {
	repo.produtos[produto.ID] = &produto
}

func (repo *RepositorioProdutoMemoria) Listar() []Produto {
	lista := make([]Produto, 0, len(repo.produtos))
	for _, produto := range repo.produtos {
		lista = append(lista, *produto)
	}
	return lista
}

// para pedidos
type RepositorioPedido interface {
	Salvar(pedido Pedido)
	BuscarPorID(ID string) (*Pedido, error)
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

func (repo *RepositorioPedidoMemoria) BuscarPorID(ID string) (*Pedido, error) {
	pedido, ok := repo.pedidos[ID]
	if !ok {
		return nil, ErrPedidoNaoEncontrado
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
