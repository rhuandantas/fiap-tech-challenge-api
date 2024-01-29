package domain

import (
	"time"
)

const (
	StatusAprovado = "aprovado"
	StatusRecusado = "recusado"
)

type Pagamento struct {
	Id        int64 `json:"id" xorm:"pk autoincr 'pagamento_id'"`
	PedidoId  int64 `xorm:"index unique"`
	Status    string
	Tipo      string
	Valor     float32
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}
