package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
)

type AtualizaStatusPedidoUC interface {
	Atualiza(ctx context.Context, status string, id int64) error
}

type atualizaStatusPedido struct {
	repo     repository.PedidoRepo
	filaRepo repository.FilaRepo
}

func (uc atualizaStatusPedido) Atualiza(ctx context.Context, status string, id int64) error {
	if err := ValidaStatuses([]string{status}); err != nil {
		return err
	}

	_, err := uc.repo.PesquisaPorID(ctx, id)
	if err != nil {
		return err
	}

	if err = uc.filaRepo.AtualizaStatus(ctx, status, id); err != nil {
		return err
	}

	return uc.repo.AtualizaStatus(ctx, status, id)
}

func NewAtualizaStatusPedidoUC(repo repository.PedidoRepo, filaRepo repository.FilaRepo) AtualizaStatusPedidoUC {
	return &atualizaStatusPedido{
		repo:     repo,
		filaRepo: filaRepo,
	}
}
