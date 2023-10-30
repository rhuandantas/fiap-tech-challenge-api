package domain

import (
	"errors"
	"strings"
	"time"
)

const (
	CategoriaLanche         = "lanche"
	CategoriaBebida         = "bebida"
	CategoriaAcompanhamento = "acompanhamento"
)

type ProdutoRequest struct {
	Descricao string `json:"descricao" validate:"required"`
	Categoria string `json:"categoria" validate:"required"`
}

type Produto struct {
	Id        int64     `json:"id" xorm:"pk autoincr 'produto_id'"`
	Descricao string    `json:"descricao" xorm:"unique"`
	Categoria string    `json:"categoria"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

func (p *Produto) ValidaCategoria() error {
	cat := strings.ToLower(p.Categoria)
	if cat == CategoriaLanche || cat == CategoriaBebida || cat == CategoriaAcompanhamento {
		return nil
	}

	return errors.New("categoria invalida")
}
