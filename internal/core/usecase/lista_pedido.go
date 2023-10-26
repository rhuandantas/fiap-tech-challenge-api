package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	"strings"
)

var validStatusesSet = map[string]bool{
	domain.StatusEmpreparacao: true,
	domain.StatusFinalizada:   true,
	domain.StatusPronto:       true,
	domain.StatusRecebido:     true,
}

type ListarPedidos interface {
	lista(ctx context.Context) ([]*domain.Pedido, error)
}

type ListarPedidoPorStatus interface {
	ListaPorStatus(ctx context.Context, statuses []string) ([]*domain.Pedido, error)
}

type listaPedidoPorStatus struct {
	repo repository.PedidoRepo
}

func (uc *listaPedidoPorStatus) ListaPorStatus(ctx context.Context, statuses []string) ([]*domain.Pedido, error) {
	if err := ValidaStatuses(statuses); err != nil {
		return nil, err
	}

	return uc.repo.PesquisaPorStatus(ctx, statuses)
}

func NewListaPedidoPorStatus(repo repository.PedidoRepo) ListarPedidoPorStatus {
	return &listaPedidoPorStatus{
		repo: repo,
	}
}

func ValidaStatuses(statuses []string) error {
	for _, s := range statuses {
		if validStatusesSet[strings.ToLower(s)] {
			return commons.BadRequest.New(fmt.Sprintf("%s não é um status valido", s))
		}
	}

	return nil
}
