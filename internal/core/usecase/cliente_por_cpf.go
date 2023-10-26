package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type PesquisarClientePorCPF interface {
	Pesquisa(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
}

type pesquisarClientePorCPFUC struct {
	clienteRepo repository.ClienteRepo
}

func NewPesquisarClientePorCpf(clienteRepo repository.ClienteRepo) PesquisarClientePorCPF {
	return &pesquisarClientePorCPFUC{
		clienteRepo: clienteRepo,
	}
}

func (uc *pesquisarClientePorCPFUC) Pesquisa(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error) {
	return uc.clienteRepo.PesquisaPorCPF(ctx, cliente)
}
