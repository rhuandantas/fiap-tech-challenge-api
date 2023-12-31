package repository

import (
	"context"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"

	"xorm.io/xorm"
)

const tableName string = "cliente"

type cliente struct {
	session *xorm.Session
}

type ClienteRepo interface {
	Insere(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
	PesquisaPorCPF(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error)
	PesquisaPorId(ctx context.Context, id int64) (*domain.Cliente, error)
}

func NewClienteRepo(connector DBConnector) ClienteRepo {
	session := connector.GetORM().Table(tableName)
	return &cliente{
		session: session,
	}
}

func (r *cliente) Insere(ctx context.Context, cliente *domain.Cliente) (*domain.Cliente, error) {
	_, err := r.session.Context(ctx).Insert(cliente)
	if err != nil {
		if commons.IsDuplicatedEntryError(err) {
			return nil, commons.BadRequest.New("cliente já existe")
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

func (r *cliente) PesquisaPorId(ctx context.Context, id int64) (*domain.Cliente, error) {
	c := domain.Cliente{
		Id: id,
	}
	found, err := r.session.Context(ctx).Get(&c)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, commons.NotFound.New("cliente não encontrado")
	}

	return &c, nil
}
