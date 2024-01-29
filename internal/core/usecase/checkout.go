package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	"github.com/joomcode/errorx"
)

type RealizarCheckout interface {
	Checkout(ctx context.Context, pedidoId int64, status string) error
}

type realizaCheckout struct {
	pedidoRepo    repository.PedidoRepo
	pagamentoRepo repository.PagamentoRepo
	fila          CadastrarFila
}

func (uc *realizaCheckout) Checkout(ctx context.Context, pedidoId int64, status string) error {
	pedidoDTO, err := uc.pedidoRepo.PesquisaPorID(ctx, pedidoId)
	if err != nil {
		return err
	}

	if pedidoDTO == nil {
		return commons.BadRequest.New(fmt.Sprint("pedido não encontrado"))
	}

	existe, err := uc.pagamentoRepo.PesquisaPorPedidoID(ctx, pedidoId)
	if err != nil {
		return err
	}

	if existe != nil {
		return errorx.IllegalState.New(fmt.Sprintf("pedido já tem um pagamento com status %s", existe.Status))
	}

	switch status {
	case domain.StatusAprovado:
		if err = uc.atualizaPagamentoAprovado(ctx, status, pedidoId); err != nil {
			return err
		}
	case domain.StatusRecusado:
		if err = uc.pagamentoRepo.Insere(ctx, &domain.Pagamento{
			Status:   domain.StatusRecusado,
			PedidoId: pedidoId,
		}); err != nil {
			return err
		}

		err = uc.pedidoRepo.AtualizaStatus(ctx, domain.StatusPagamentoRecusado, pedidoId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc *realizaCheckout) atualizaPagamentoAprovado(ctx context.Context, status string, pedidoID int64) error {
	if err := uc.pagamentoRepo.Insere(ctx, &domain.Pagamento{
		Status:   domain.StatusAprovado,
		PedidoId: pedidoID,
	}); err != nil {
		return err
	}

	if err := uc.fila.Cadastra(ctx, &domain.Fila{Status: domain.StatusPagamentoAprovado, PedidoId: pedidoID}); err != nil {
		return err
	}

	if err := uc.pedidoRepo.AtualizaStatus(ctx, domain.StatusPagamentoAprovado, pedidoID); err != nil {
		return err
	}

	return nil
}

func NewRealizaCheckout(pedidoRepo repository.PedidoRepo, pagamentoRepo repository.PagamentoRepo, fila CadastrarFila) RealizarCheckout {
	return &realizaCheckout{
		pedidoRepo:    pedidoRepo,
		pagamentoRepo: pagamentoRepo,
		fila:          fila,
	}
}
