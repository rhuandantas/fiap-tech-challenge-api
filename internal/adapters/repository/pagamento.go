package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/joomcode/errorx"
	"xorm.io/xorm"
)

const tableNamePagamento string = "pagamento"

type pagamento struct {
	session *xorm.Session
}

type PagamentoRepo interface {
	Insere(ctx context.Context, pagamento *domain.Pagamento) error
	AtualizaStatus(ctx context.Context, status string, pedidoId int64) error
	PesquisaPorPedidoID(ctx context.Context, pedidoId int64) (*domain.Pagamento, error)
}

func NewPagamentoRepo(connector DBConnector) PagamentoRepo {
	session := connector.GetORM().Table(tableNamePagamento)
	return &pagamento{
		session: session,
	}
}

func (f *pagamento) Insere(ctx context.Context, pagamento *domain.Pagamento) error {
	_, err := f.session.Context(ctx).Insert(pagamento)
	if err != nil {
		return err
	}

	return nil
}

func (f *pagamento) AtualizaStatus(ctx context.Context, status string, pedidoId int64) error {
	_, err := f.session.Context(ctx).Where("pedido_id = ?", pedidoId).Update(&domain.Pagamento{Status: status})
	if err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}

func (f *pagamento) PesquisaPorPedidoID(ctx context.Context, pedidoId int64) (*domain.Pagamento, error) {
	dto := &domain.Pagamento{PedidoId: pedidoId}
	has, err := f.session.Context(ctx).Get(dto)
	if err != nil {
		return nil, err
	}

	if has {
		return dto, nil
	}

	return nil, nil
}
