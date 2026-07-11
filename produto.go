package main

type Produto struct {
	Id      string
	Nome    string
	Preco   float64
	Estoque int
}

func (item *Produto) ReduzirEstoque(quantidade int) error {
	if quantidade > item.Estoque {
		return ErrEstoqueInsuficiente
	}
	item.Estoque = item.Estoque - quantidade
	return nil
}

func (item *Produto) AdicionaProduto(quantidade int) {
	item.Estoque = item.Estoque + quantidade
}
