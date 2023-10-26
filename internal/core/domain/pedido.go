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

type Pedido struct {
	Id        int64     `json:"id"`
	Status    string    `xorm:"status"`
	Cliente   Cliente   `xorm:"extends"`
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}

type Status struct {
	Id        int64
	IdPedido  int64
	Name      string    `xorm:"status"`
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}
