package domain

import (
	"github.com/google/uuid"
	"time"
)

type Pedido struct {
	Id        uuid.UUID `xorm:"pk varchar(64)"`
	Status    []Status  `xorm:"extends"`
	Cliente   Cliente   `xorm:"extends"`
	Produtos  []Produto
	Pagamento Pagamento `xorm:"extends"`
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}

type Status struct {
	Id        uuid.UUID `xorm:"pk varchar(64)"`
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}
