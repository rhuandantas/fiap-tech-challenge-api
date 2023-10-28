package domain

import (
	"bytes"
	"errors"
	"fiap-tech-challenge-api/internal/core/commons"
	"github.com/paemuri/brdoc"
	"time"
	"unicode"
)

type Cliente struct {
	Id              int64     `json:"id" xorm:"pk autoincr 'cliente_id'"`
	Cpf             string    `json:"cpf" validate:"required" xorm:"unique"`
	Nome            string    `json:"nome" validate:"required"`
	Email           string    `json:"email" validate:"email"`
	DataAniversario time.Time `json:"data_aniversario"`
	Telefone        string    `json:"telefone"`
	CreatedAt       time.Time `json:"created_at" xorm:"created"`
	UpdatedAt       time.Time `json:"updated_at" xorm:"updated"`
}

func (c *Cliente) ValidateCPF() error {
	if !brdoc.IsCPF(c.Cpf) {
		return errors.New(commons.CpfInvalido)
	}

	c.limpaCaracteresEspeciais()

	return nil
}

func (c *Cliente) limpaCaracteresEspeciais() {
	buf := bytes.NewBufferString("")
	for _, r := range c.Cpf {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	c.Cpf = buf.String()
}