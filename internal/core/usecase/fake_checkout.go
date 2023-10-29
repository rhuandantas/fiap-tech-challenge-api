package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type RealizarCheckout interface {
	FakeCheckout(ctx context.Context, pedidoId int64) error
}

type realizaCheckout struct {
	pedidoRepo repository.PedidoRepo
	filaRepo   repository.FilaRepo
}

func (uc *realizaCheckout) FakeCheckout(ctx context.Context, pedidoId int64) error {
	_, err := uc.pedidoRepo.PesquisaPorID(ctx, pedidoId)
	if err != nil {
		return err
	}

	err = uc.pedidoRepo.AtualizaStatus(ctx, domain.StatusRecebido, pedidoId)
	if err != nil {
		return err
	}

	return uc.filaRepo.AtualizaStatus(ctx, domain.StatusRecebido, pedidoId)
}

func NewRealizaCheckout(pedidoRepo repository.PedidoRepo, filaRepo repository.FilaRepo) RealizarCheckout {
	return &realizaCheckout{
		pedidoRepo: pedidoRepo,
		filaRepo:   filaRepo,
	}
}
