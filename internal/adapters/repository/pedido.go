package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/joomcode/errorx"
	"xorm.io/xorm"
)

const tableNamePedido string = "pedido"

type pedido struct {
	session *xorm.Session
}

type PedidoRepo interface {
	Insere(ctx context.Context, pedido *domain.PedidoDTO) (*domain.PedidoDTO, error)
	PesquisaPorStatus(ctx context.Context, statuses []string) ([]*domain.PedidoDTO, error)
	Atualiza(ctx context.Context, pedido *domain.PedidoDTO, id int64) error
}

func NewPedidoRepo(connector DBConnector) PedidoRepo {
	session := connector.GetORM().Table(tableNamePedido)
	return &pedido{
		session: session,
	}
}

func (p *pedido) Insere(ctx context.Context, pedido *domain.PedidoDTO) (*domain.PedidoDTO, error) {
	pedido.Status = domain.StatusRecebido
	if _, err := p.session.Context(ctx).Insert(pedido); err != nil {
		return nil, err
	}

	return pedido, nil
}

func (p *pedido) PesquisaPorStatus(ctx context.Context, statuses []string) ([]*domain.PedidoDTO, error) {
	pedidos := make([]*domain.PedidoDTO, 0)
	err := p.session.Context(ctx).In("status", statuses).Find(&pedidos)
	if err != nil {
		return nil, errorx.InternalError.New(err.Error())
	}

	return pedidos, nil
}

func (p *pedido) Atualiza(ctx context.Context, new *domain.PedidoDTO, id int64) error {
	_, err := p.session.Context(ctx).ID(id).Update(new)
	if err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}
