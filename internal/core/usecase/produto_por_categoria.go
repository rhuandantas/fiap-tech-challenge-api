package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
)

type PegarProdutoPorCategoria interface {
	Pega(ctx context.Context, p *domain.Produto) ([]*domain.Produto, error)
}

type pegaProdutoPorCategoria struct {
	repo repository.ProdutoRepo
}

func (uc pegaProdutoPorCategoria) Pega(ctx context.Context, p *domain.Produto) ([]*domain.Produto, error) {
	return uc.repo.PesquisaPorCategoria(ctx, p)
}

func NewPegaProdutoPorCategoria(repo repository.ProdutoRepo) PegarProdutoPorCategoria {
	return &pegaProdutoPorCategoria{
		repo: repo,
	}
}
