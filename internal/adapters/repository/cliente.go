package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/google/uuid"
	"xorm.io/xorm"
)

const tableName string = "cliente"

type cliente struct {
	session *xorm.Session
}

type ClienteRepo interface {
	Insere(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
	PesquisaPorCPF(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
}

func NewClienteRepo(connector DBConnector) ClienteRepo {
	session := connector.GetORM().Table(tableName)
	return &cliente{
		session: session,
	}
}

func (r *cliente) Insere(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error) {
	newUUID, _ := uuid.NewUUID()
	cliente.Id = newUUID
	_, err := r.session.Context(ctx).Insert(cliente)
	if err != nil {
		if commons.IsDuplicatedEntryError(err) {
			return nil, commons.BadRequest.New("client já existe")
		}

		return nil, err
	}

	return cliente, nil
}

func (r *cliente) PesquisaPorCPF(ctx context.Context, c *domain.Cliente) (*domain.Cliente, error) {
	cliente := domain.Cliente{
		Cpf: c.Cpf,
	}
	found, err := r.session.Context(ctx).Get(&cliente)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, commons.NotFound.New("cliente não encontrado")
	}

	return &cliente, nil
}
