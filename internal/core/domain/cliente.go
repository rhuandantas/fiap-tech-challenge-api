package domain

import (
	"bytes"
	"errors"
	"fiap-tech-challenge-api/internal/core/commons"
	"time"
	"unicode"

	"github.com/paemuri/brdoc"
)

type ClienteRequest struct {
	Cpf      string `json:"cpf" validate:"required" xorm:"unique"`
	Nome     string `json:"nome" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Telefone string `json:"telefone"`
}

type LGPDClienteRequest struct {
	Cpf      string `json:"cpf" validate:"required" xorm:"unique"`
}

type Cliente struct {
	Id        int64     `json:"id" xorm:"pk autoincr 'cliente_id'"`
	Cpf       *string    `json:"cpf" xorm:"null unique"`
	Nome      *string    `json:"nome" xorm:"null"`
	Email     *string    `json:"email" xorm:"null"`
	Telefone  *string    `json:"telefone" xorm:"null"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func (c *Cliente) ValidateCPF() error {
	if !brdoc.IsCPF(DerefString(c.Cpf)) {
		return errors.New(commons.CpfInvalido)
	}

	c.limpaCaracteresEspeciais()

	return nil
}

func DerefString(s *string) string {
    if s != nil {
        return *s
    }

    return ""
}


func (c *Cliente) limpaCaracteresEspeciais() {
	buf := bytes.NewBufferString("")
	for _, r := range DerefString(c.Cpf) {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	varAux := string(buf.String())
	c.Cpf = &varAux
}
