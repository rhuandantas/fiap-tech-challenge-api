package domain

import (
	"time"
)

const (
	StatusRecebido     string = "recebido"
	StatusEmpreparacao string = "em_preparacao"
	StatusPronto       string = "pronto"
	StatusFinalizada   string = "finalizada"
)

type PedidoRequest struct {
	ClienteId  int64   `json:"cliente_id" validate:"required"`
	ProdutoIds []int64 `json:"produtos" validate:"required"`
	Observacao string  `json:"observacao"`
}

type PedidoDTO struct {
	Id         int64      `xorm:"pk autoincr 'pedido_id'"`
	ClienteId  int64      `xorm:"index"`
	Cliente    *Cliente   `xorm:"-"`
	Produtos   []*Produto `xorm:"-"`
	Status     string
	Observacao string
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type PedidoProduto struct {
	Id        int64
	PedidoId  int64     `xorm:"index"`
	ProdutoId int64     `xorm:"index"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
}

func (dto *PedidoDTO) TableName() string {
	return "pedido"
}

type PedidoResponse struct {
	*Pedido
}

type Pedido struct {
	Id         int64      `json:"id"`
	Status     string     `json:"status"`
	Cliente    *Cliente   `json:"cliente,omitempty"`
	Produtos   []*Produto `json:"produtos,omitempty"`
	Observacao string     `json:"observacao"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated"`
}

type Status struct {
	Id        int64
	PedidoId  int64     `xorm:"index"`
	Status    string    `xorm:"status"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
