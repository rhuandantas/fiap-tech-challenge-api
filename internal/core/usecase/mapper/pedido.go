package mapper

import (
	"fiap-tech-challenge-api/internal/core/domain"
)

type Pedido interface {
	MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO
	MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido
	MapDTOToModel(req *domain.PedidoDTO) *domain.Pedido
	MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse
}

type pedido struct {
}

func (p pedido) MapDTOToResponse(dto *domain.PedidoDTO) *domain.PedidoResponse {
	return &domain.PedidoResponse{
		Pedido: &domain.Pedido{
			Id:         dto.Id,
			Status:     dto.Status,
			Cliente:    dto.Cliente,
			Produtos:   dto.Produtos,
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		},
	}
}

func (p pedido) MapReqToDTO(req *domain.PedidoRequest) *domain.PedidoDTO {
	return &domain.PedidoDTO{
		ClienteId:  req.ClienteId,
		Observacao: req.Observacao,
	}
}

func (p pedido) MapDTOToModel(req *domain.PedidoDTO) *domain.Pedido {
	//TODO implement me
	panic("implement me")
}

func (p pedido) MapDTOToModels(req []*domain.PedidoDTO) []*domain.Pedido {
	pedidos := make([]*domain.Pedido, len(req))
	for i, dto := range req {
		pedidos[i] = &domain.Pedido{
			Id:         dto.Id,
			Status:     dto.Status,
			Cliente:    dto.Cliente,
			Produtos:   dto.Produtos,
			Observacao: dto.Observacao,
			CreatedAt:  dto.CreatedAt,
			UpdatedAt:  dto.UpdatedAt,
		}
	}

	return pedidos
}

func NewPedidoMapper() Pedido {
	return &pedido{}
}
