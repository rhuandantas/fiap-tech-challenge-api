package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	"github.com/joomcode/errorx"
)

type PesquisaPagamento interface {
	PesquisaPorPedidoID(ctx context.Context, pedidoId int64) (*domain.Pagamento, error)
}

type pesquisaPagamento struct {
	pedidoRepo    repository.PedidoRepo
	pagamentoRepo repository.PagamentoRepo
}

func (uc *pesquisaPagamento) PesquisaPorPedidoID(ctx context.Context, pedidoId int64) (*domain.Pagamento, error) {
	pedidoDTO, err := uc.pedidoRepo.PesquisaPorID(ctx, pedidoId)
	if err != nil {
		return nil, err
	}

	if pedidoDTO == nil {
		return nil, commons.BadRequest.New(fmt.Sprint("pedido não encontrado"))
	}

	existe, err := uc.pagamentoRepo.PesquisaPorPedidoID(ctx, pedidoId)
	if err != nil {
		return nil, err
	}

	if existe == nil {
		return nil, errorx.IllegalState.New(fmt.Sprintf("nenhum pagamento encontrado para pedido %d", pedidoId))
	}

	return existe, nil
}

func NewPesquisaPagamento(pedidoRepo repository.PedidoRepo, pagamentoRepo repository.PagamentoRepo) PesquisaPagamento {
	return &pesquisaPagamento{
		pedidoRepo:    pedidoRepo,
		pagamentoRepo: pagamentoRepo,
	}
}
