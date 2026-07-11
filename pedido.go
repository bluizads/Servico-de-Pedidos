package main

type StatusPedido string

const (
	StatusPendente  StatusPedido = "PENDENTE"
	StatusPago      StatusPedido = "PAGO"
	StatusCancelado StatusPedido = "CANCELADO"
)

type ItemPedido struct {
	ProdutoId     string
	PrecoNaCompra float64
	Quantidade    int
}

type Pedido struct {
	ID      string
	Cliente string
	Itens   []ItemPedido
	Status  StatusPedido
}

func (compra Pedido) CalcularTotal() float64 {
	total := 0.0
	for _, item := range compra.Itens {
		total += item.PrecoNaCompra * float64(item.Quantidade)
	}
	return total
}

// recebe o ponteiro
func (compra *Pedido) Pagar() error {
	if compra.Status != StatusPendente {
		return ErrMudancasStatusInvalida
	}
	compra.Status = StatusPago
	return nil
}
