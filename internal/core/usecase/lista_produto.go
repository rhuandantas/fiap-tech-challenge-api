package usecase

import (
	"context"
	"errors"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
	"fmt"
	serverErr "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
	"strconv"
)

type PegaPorIDS interface {
	PegaPorIDS(ctx context.Context, ids []string) ([]*domain.Produto, error)
}

type pegaPorIDS struct {
	repo repository.ProdutoRepo
}

func (uc pegaPorIDS) PegaPorIDS(ctx context.Context, ids []string) ([]*domain.Produto, error) {
	intIDS, err := uc.stringToIDS(ids)
	if err != nil {
		return nil, serverErr.BadRequest.New("ids inválidos")
	}

	return uc.repo.PesquisaPorIDS(ctx, intIDS)
}

func (uc pegaPorIDS) stringToIDS(ids []string) ([]int64, error) {
	intIDS := make([]int64, len(ids))
	for _, id := range ids {
		intID, err := strconv.Atoi(id)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("produto id %s inválido", id))
		}
		intIDS = append(intIDS, int64(intID))
	}

	return intIDS, nil
}

func NewPegaPorIDS(repo repository.ProdutoRepo) PegaPorIDS {
	return &pegaPorIDS{
		repo: repo,
	}
}
