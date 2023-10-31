package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
)

type CadastrarPedido interface {
	Cadastra(ctx context.Context, pedido *domain.PedidoRequest) (*domain.PedidoResponse, error)
}

type cadastraPedido struct {
	repo         repository.PedidoRepo
	pedidoRepo   repository.PedidoProdutoRepo
	clienteRepo  repository.ClienteRepo
	prodRepo     repository.ProdutoRepo
	filaRepo     repository.FilaRepo
	mapperPedido mapper.Pedido
}

func (uc cadastraPedido) Cadastra(ctx context.Context, req *domain.PedidoRequest) (*domain.PedidoResponse, error) {
	ids, err := uc.prodRepo.PesquisaPorIDS(ctx, req.ProdutoIds)
	if err != nil {
		return nil, err
	}

	if len(ids) != len(req.ProdutoIds) {
		return nil, commons.BadRequest.New("pedido contém produto(s) inválido(s)")
	}

	dto, err := uc.repo.Insere(ctx, uc.mapperPedido.MapReqToDTO(req))
	if err != nil {
		return nil, err
	}
	pedidoProdutos := make([]*domain.PedidoProduto, len(req.ProdutoIds))
	for i, id := range req.ProdutoIds {
		pedidoProdutos[i] = &domain.PedidoProduto{
			PedidoId:  dto.Id,
			ProdutoId: id,
		}
	}

	if err = uc.pedidoRepo.Insere(ctx, pedidoProdutos); err != nil {
		return nil, err
	}

	if err = uc.filaRepo.Insere(ctx, &domain.Fila{
		Status:     domain.StatusRecebido,
		PedidoId:   dto.Id,
		Observacao: dto.Observacao,
	}); err != nil {
		return nil, err
	}

	return uc.mapperPedido.MapDTOToResponse(dto), err
}

func NewCadastraPedido(repo repository.PedidoRepo, pedidoRepo repository.PedidoProdutoRepo, mapperPedido mapper.Pedido, clienteRepo repository.ClienteRepo,
	prodRepo repository.ProdutoRepo, filaRepo repository.FilaRepo) CadastrarPedido {
	return &cadastraPedido{
		repo:         repo,
		mapperPedido: mapperPedido,
		pedidoRepo:   pedidoRepo,
		prodRepo:     prodRepo,
		clienteRepo:  clienteRepo,
		filaRepo:     filaRepo,
	}
}
