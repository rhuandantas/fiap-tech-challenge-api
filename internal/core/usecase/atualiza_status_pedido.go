package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
)

type AtualizaStatusPedidoUC interface {
	Atualiza(ctx context.Context, status string, id int64) error
}

type atualizaStatusPedido struct {
	repo repository.PedidoRepo
}

func (uc atualizaStatusPedido) Atualiza(ctx context.Context, status string, id int64) error {
	return uc.repo.AtualizaStatus(ctx, status, id)
}

func NewAtualizaStatusPedidoUC(repo repository.PedidoRepo) AtualizaStatusPedidoUC {
	return &atualizaStatusPedido{
		repo: repo,
	}
}
