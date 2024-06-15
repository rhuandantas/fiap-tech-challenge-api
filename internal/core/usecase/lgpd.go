package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type LGPD interface {
	Anonimizar(ctx context.Context, cliente *domain.Cliente) error
}

type anonimizarClienteUC struct {
	clienteRepo repository.ClienteRepo
}

func NewLGPD(clienteRepo repository.ClienteRepo) LGPD {
	return &anonimizarClienteUC{
		clienteRepo: clienteRepo,
	}
}


func (uc *anonimizarClienteUC) Anonimizar(ctx context.Context, cliente *domain.Cliente) error {
	return uc.clienteRepo.Anonimizar(ctx, cliente)
}
