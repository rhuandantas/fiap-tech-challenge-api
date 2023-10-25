package domain

import (
	"github.com/google/uuid"
	"time"
)

type Pagamento struct {
	Id        uuid.UUID `xorm:"pk varchar(64)"`
	Tipo      string
	Valor     float32
	Status    string
	CreatedAt time.Time `xorm:"created"`
	Update    time.Time `xorm:"updated"`
}
