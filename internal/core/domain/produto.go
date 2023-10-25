package domain

import (
	"github.com/google/uuid"
	"time"
)

type Produto struct {
	Id        uuid.UUID `xorm:"pk varchar(64)"`
	Descricao string    `xorm:"unique"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
