package domain

import (
	"time"
)

const (
	PagamentoTipoDinheiro = "dinheiro"
	PagamentoTipoCredito  = "credito"
	PagamentoTipoDebito   = "debito"
)

type Pagamento struct {
	Id        int64 `json:"id" xorm:"pk autoincr 'pagamento_id'"`
	Tipo      string
	Valor     float32
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}
