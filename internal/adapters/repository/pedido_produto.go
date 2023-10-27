package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/joomcode/errorx"

	"xorm.io/xorm"
)

type pedidoProduto struct {
	session *xorm.Session
}

type PedidoProdutoRepo interface {
	Insere(ctx context.Context, pedidoProdutos []*domain.PedidoProduto) error
	PesquisaPedidoProduto(ctx context.Context, pedidoProduto *domain.PedidoProduto) ([]*domain.Produto, error)
}

func NewPedidoProdutoRepo(connector DBConnector) PedidoProdutoRepo {
	session := connector.GetORM().Table("pedido_produto")
	return &pedidoProduto{
		session: session,
	}
}

func (p *pedidoProduto) Insere(ctx context.Context, pedidoProdutos []*domain.PedidoProduto) error {
	if _, err := p.session.Context(ctx).Insert(pedidoProdutos); err != nil {
		return errorx.InternalError.New(err.Error())
	}

	return nil
}

func (p *pedidoProduto) PesquisaPedidoProduto(ctx context.Context, filter *domain.PedidoProduto) ([]*domain.Produto, error) {
	produtos := make([]*domain.Produto, 0)
	err := p.session.Context(ctx).
		Join("INNER", "produto", "pedido_produto.produto_id = produto.produto_id").
		Join("INNER", "pedido", "pedido.pedido_id = pedido_produto.pedido_id").
		Find(&produtos, filter)
	if err != nil {
		return nil, err
	}

	return produtos, nil
}
