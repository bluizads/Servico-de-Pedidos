package main

import "errors"

var (
	ErrProdutoNaoEncontrado   = errors.New("Produto nao encontrado")
	ErrPedidoNaoEcncontrado   = errors.New("pedido nao encontrado")
	ErrQuantidadeInvalida     = errors.New("Quantidade invalida")
	ErrEstoqueInsuficiente    = errors.New("estoque insuficiente")
	ErrClienteInvalido        = errors.New("cliente inválido")
	ErrPedidoVazio            = errors.New("Pedido vazio")
	ErrMudancasStatusInvalida = errors.New("mudanca de status invalida")
)
