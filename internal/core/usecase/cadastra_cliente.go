package usecase

import (
	"context"
	"errors"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
	_error "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
)

type CadastrarClienteUseCase interface {
	Cadastra(ctx context.Context, cliente *domain.ClienteRequest) (*domain.Cliente, error)
}

type cadastraClienteUC struct {
	clienteRepo repository.ClienteRepo
}

func NewCadastraCliente(clienteRepo repository.ClienteRepo) CadastrarClienteUseCase {
	return &cadastraClienteUC{
		clienteRepo: clienteRepo,
	}
}

func (uc *cadastraClienteUC) Cadastra(ctx context.Context, cliente *domain.ClienteRequest) (*domain.Cliente, error) {
	found, err := uc.clienteRepo.PesquisaPorCPF(ctx, domain.NewClient(cliente))
	if err != nil {
		if errors.Is(err, _error.BadRequest.New("cliente não encontrado")) {
			return nil, err
		}
	}

	if found != nil {
		return found, _error.BadRequest.New("cliente já existe")
	}

	return uc.clienteRepo.Insere(ctx, domain.NewClient(cliente))
}
