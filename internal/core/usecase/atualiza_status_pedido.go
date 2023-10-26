package usecase

type AtualizaStatusPedidoUC interface {
	atualizar()
}

type atualizaStatusPedido struct {
}

func NewAtualizaStatusPedidoUC() AtualizaStatusPedidoUC {
	return nil
}
