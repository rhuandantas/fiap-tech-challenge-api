package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type CadastrarProduto interface {
	Cadastra(ctx context.Context, produto *domain.Produto) (*domain.Produto, error)
}

type cadastraProduto struct {
	repo repository.ProdutoRepo
}

func (uc cadastraProduto) Cadastra(ctx context.Context, produto *domain.Produto) (*domain.Produto, error) {
	return uc.repo.Insere(ctx, produto)
}

func NewCadastraProduto(repo repository.ProdutoRepo) CadastrarProduto {
	return &cadastraProduto{
		repo: repo,
	}
}
