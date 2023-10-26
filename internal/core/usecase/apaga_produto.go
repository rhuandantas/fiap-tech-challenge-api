package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type ApagarProduto interface {
	Apaga(ctx context.Context, produto *domain.Produto) error
}

type apagaProduto struct {
	repo repository.ProdutoRepo
}

func (a apagaProduto) Apaga(ctx context.Context, produto *domain.Produto) error {
	return a.repo.Apaga(ctx, produto)
}

func NewApagaProduto(repo repository.ProdutoRepo) ApagarProduto {
	return &apagaProduto{
		repo: repo,
	}
}
