package domain

import (
	"github.com/google/uuid"
	"time"
)

type Cliente struct {
	Id              uuid.UUID `xorm:"pk varchar(64)"`
	Cpf             string    `xorm:"unique"`
	Nome            string
	Email           string
	DataAniversario time.Time
	Telefone        string
	CreatedAt       time.Time `xorm:"created"`
	UpdatedAt       time.Time `xorm:"updated"`
}
