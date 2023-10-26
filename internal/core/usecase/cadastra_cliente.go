package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type CadastrarClienteUseCase interface {
	Cadastra(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
}

type cadastraClienteUC struct {
	clienteRepo repository.ClienteRepo
}

func NewCadastraCliente(clienteRepo repository.ClienteRepo) CadastrarClienteUseCase {
	return &cadastraClienteUC{
		clienteRepo: clienteRepo,
	}
}

func (uc *cadastraClienteUC) Cadastra(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error) {
	return uc.clienteRepo.Insere(ctx, cliente)
}
