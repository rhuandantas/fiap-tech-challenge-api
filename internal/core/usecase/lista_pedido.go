package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
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
	repo         repository.PedidoRepo
	mapperPedido mapper.Pedido
}

func (uc *listaPedidoPorStatus) ListaPorStatus(ctx context.Context, statuses []string) ([]*domain.Pedido, error) {
	if err := ValidaStatuses(statuses); err != nil {
		return nil, err
	}

	pedidos, err := uc.repo.PesquisaPorStatus(ctx, statuses)
	if err != nil {
		return nil, err
	}
	return uc.mapperPedido.MapDTOToModels(pedidos), nil
}

func NewListaPedidoPorStatus(repo repository.PedidoRepo, mapperPedido mapper.Pedido) ListarPedidoPorStatus {
	return &listaPedidoPorStatus{
		repo:         repo,
		mapperPedido: mapperPedido,
	}
}

func ValidaStatuses(statuses []string) error {
	for _, s := range statuses {
		if !validStatusesSet[strings.ToLower(s)] {
			return commons.BadRequest.New(fmt.Sprintf("%s não é um status valido", s))
		}
	}

	return nil
}
