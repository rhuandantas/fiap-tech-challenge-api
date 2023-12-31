package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
	"strconv"
	"strings"
)

type PegarDetalhePedido interface {
	Pesquisa(ctx context.Context, id int64) (*domain.PedidoResponse, error)
}

type pegaDetalhePedido struct {
	repo         repository.PedidoRepo
	repoPedProd  repository.PedidoProdutoRepo
	repoProd     repository.ProdutoRepo
	repoCli      repository.ClienteRepo
	mapperPedido mapper.Pedido
}

func (uc *pegaDetalhePedido) Pesquisa(ctx context.Context, id int64) (*domain.PedidoResponse, error) {
	dto, err := uc.repo.PesquisaPorID(ctx, id)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	for _, prodId := range strings.Split(dto.ProdutoIDS, ",") {
		idInt, _ := strconv.Atoi(prodId)
		ids = append(ids, int64(idInt))
	}
	produtos, err := uc.repoProd.PesquisaPorIDS(ctx, ids)
	if err != nil {
		return nil, err
	}

	dto.Produtos = produtos

	return uc.mapperPedido.MapDTOToResponse(dto), nil
}

func NewPegaDetalhePedido(repo repository.PedidoRepo,
	repoPedProd repository.PedidoProdutoRepo,
	repoProd repository.ProdutoRepo,
	repoCli repository.ClienteRepo,
	mapperPedido mapper.Pedido,
) PegarDetalhePedido {
	return &pegaDetalhePedido{
		repo:         repo,
		repoPedProd:  repoPedProd,
		repoProd:     repoProd,
		repoCli:      repoCli,
		mapperPedido: mapperPedido,
	}
}
