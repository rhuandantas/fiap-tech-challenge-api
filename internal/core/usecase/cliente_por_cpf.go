package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type PesquisarCliente interface {
	PesquisaPorCPF(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
	PesquisaPorID(ctx context.Context, id int64) (*domain.Cliente, error)
}

type pesquisarClienteUC struct {
	clienteRepo repository.ClienteRepo
}

func NewPesquisarCliente(clienteRepo repository.ClienteRepo) PesquisarCliente {
	return &pesquisarClienteUC{
		clienteRepo: clienteRepo,
	}
}

func (uc *pesquisarClienteUC) PesquisaPorCPF(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error) {
	return uc.clienteRepo.PesquisaPorCPF(ctx, cliente)
}
func (uc *pesquisarClienteUC) PesquisaPorID(ctx context.Context, id int64) (*domain.Cliente, error) {
	return uc.clienteRepo.PesquisaPorId(ctx, id)
}
