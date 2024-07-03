package domain

import (
	"bytes"
	"database/sql"
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
type ClienteResponse struct {
	Id        int64     `json:"id"`
	Cpf       string    `json:"cpf" validate:"required" xorm:"unique"`
	Nome      string    `json:"nome" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Telefone  string    `json:"telefone"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

type LGPDClienteRequest struct {
	Cpf string `json:"cpf" validate:"required" xorm:"unique"`
}

type Cliente struct {
	Id        int64          `json:"id" xorm:"pk autoincr 'cliente_id'"`
	Cpf       sql.NullString `json:"cpf" xorm:"null"`
	Nome      sql.NullString `json:"nome" xorm:"null"`
	Email     sql.NullString `json:"email" xorm:"null"`
	Telefone  sql.NullString `json:"telefone" xorm:"null"`
	CreatedAt time.Time      `json:"created_at" xorm:"created"`
	UpdatedAt time.Time      `json:"updated_at" xorm:"updated"`
}

func NewClient(cli *ClienteRequest) *Cliente {
	return &Cliente{
		Cpf:      setNullString(cli.Cpf),
		Nome:     setNullString(cli.Nome),
		Email:    setNullString(cli.Email),
		Telefone: setNullString(cli.Telefone),
	}
}

func NewClienteResponse(cli *Cliente) *ClienteResponse {
	return &ClienteResponse{
		Id:        cli.Id,
		Cpf:       cli.Cpf.String,
		Nome:      cli.Nome.String,
		Email:     cli.Email.String,
		Telefone:  cli.Telefone.String,
		CreatedAt: cli.CreatedAt,
		UpdatedAt: cli.UpdatedAt,
	}
}

func setNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{}
	}

	return sql.NullString{String: str, Valid: true}
}

func (c *ClienteRequest) ValidateCPF() error {
	if !brdoc.IsCPF(c.Cpf) {
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

func (c *ClienteRequest) limpaCaracteresEspeciais() {
	buf := bytes.NewBufferString("")
	for _, r := range c.Cpf {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	c.Cpf = buf.String()
}
