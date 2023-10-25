package repository

import (
	"context"
	"errors"
	"fiap-tech-challenge-api/internal/adpters/database"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"github.com/google/uuid"
	"xorm.io/xorm"
)

const tableName string = "cliente"

type ClienteRepo interface {
	Insere(ctx context.Context, cliente *domain.Cliente) error
	Lista(ctx context.Context) ([]*domain.Cliente, error)
}

type cliente struct {
	session *xorm.Session
}

func NewClienteRepo(connector database.DBConnector) ClienteRepo {
	session := connector.GetORM().Table(tableName)
	return &cliente{
		session: session,
	}
}

func (c *cliente) Insere(ctx context.Context, cliente *domain.Cliente) error {
	newUUID, _ := uuid.NewUUID()
	cliente.Id = newUUID
	_, err := c.session.Context(ctx).Insert(cliente)
	if err != nil {
		if commons.DuplicatedEntryError(err) {
			return errors.New("client already exists")
		}

		return err
	}

	return nil
}

func (c *cliente) Lista(ctx context.Context) ([]*domain.Cliente, error) {
	var clientes []*domain.Cliente
	err := c.session.Context(ctx).Find(&clientes)
	if err != nil {
		return nil, err
	}

	return clientes, nil
}
