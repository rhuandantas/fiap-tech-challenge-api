package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type LGPD interface {
	Anonimizar(ctx context.Context, cliente *domain.ClienteRequest) error
}

type anonimizarClienteUC struct {
	clienteRepo repository.ClienteRepo
}

func NewLGPD(clienteRepo repository.ClienteRepo) LGPD {
	return &anonimizarClienteUC{
		clienteRepo: clienteRepo,
	}
}

func (uc *anonimizarClienteUC) Anonimizar(ctx context.Context, cliente *domain.ClienteRequest) error {
	dto := domain.NewClient(cliente)

	cli, err := uc.clienteRepo.PesquisaPorCPF(ctx, dto)
	if err != nil {
		return err
	}

	cli.Email.String = ""
	cli.Cpf.String = ""
	cli.Nome.String = ""
	cli.Telefone.String = ""

	return uc.clienteRepo.Anonimizar(ctx, cli)
}
