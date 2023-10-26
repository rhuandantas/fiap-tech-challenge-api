package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type AtualizarProduto interface {
	Atualiza(ctx context.Context, produto *domain.Produto, id int64) error
}

type atualizaProduto struct {
	repo repository.ProdutoRepo
}

func (a atualizaProduto) Atualiza(ctx context.Context, produto *domain.Produto, id int64) error {
	return a.repo.Atualiza(ctx, produto, id)
}

func NewAtualizaProduto(repo repository.ProdutoRepo) AtualizarProduto {
	return &atualizaProduto{
		repo: repo,
	}
}
