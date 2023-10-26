package handlers

import (
	serverErr "fiap-tech-challenge-api/internal/adapters/http/error"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type Produto struct {
	validator                  util.Validator
	cadastraProdutoUC          usecase.CadastrarProduto
	pegarProdutoPorCategoriaUC usecase.PegarProdutoPorCategoria
	apagarProduto              usecase.ApagarProduto
	atualizarProduto           usecase.AtualizarProduto
}

func NewProduto(validator util.Validator,
	cadastraProdutoUC usecase.CadastrarProduto,
	pegarProdutoPorCategoriaUC usecase.PegarProdutoPorCategoria,
	apagarProduto usecase.ApagarProduto,
	atualizarProduto usecase.AtualizarProduto,
) *Produto {
	return &Produto{
		validator:                  validator,
		cadastraProdutoUC:          cadastraProdutoUC,
		pegarProdutoPorCategoriaUC: pegarProdutoPorCategoriaUC,
		apagarProduto:              apagarProduto,
		atualizarProduto:           atualizarProduto,
	}
}

func (h *Produto) RegistraRotasProduto(server *echo.Echo) {
	server.POST("/produto", h.cadastra)
	server.GET("/produtos/:categoria", h.listaPorCategoria)
	server.DELETE("/produto/:id", h.apaga)
	server.PUT("/produto/:id", h.atualiza)
}

// cadastra godoc
// @Summary cadastra um novo produto
// @Tags Produto
// @Accept json
// @Produce json
// @Success 200 {object} domain.Produto
// @Router /produto [post]
func (h *Produto) cadastra(ctx echo.Context) error {
	var (
		produto domain.Produto
		err     error
	)

	if err = ctx.Bind(&produto); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validateProdutoBody(&produto); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	newProduto, err := h.cadastraProdutoUC.Cadastra(ctx.Request().Context(), &produto)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusCreated, newProduto)
}

// listaPorCategoria godoc
// @Summary pega produtos por categoria
// @Tags Produto
// @Produce json
// @Param        categoria   path      string  true  "categoria do produto"
// @Success 200 {array} domain.Produto
// @Router /produtos/{categoria} [get]
func (h *Produto) listaPorCategoria(ctx echo.Context) error {
	categoria := ctx.Param("categoria")
	p := &domain.Produto{
		Categoria: categoria,
	}
	if err := p.ValidaCategoria(); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	produtos, err := h.pegarProdutoPorCategoriaUC.Pega(ctx.Request().Context(), p)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, produtos)
}

func (h *Produto) validateProdutoBody(p *domain.Produto) error {
	if err := h.validator.ValidateStruct(p); err != nil {
		return err
	}

	if err := p.ValidaCategoria(); err != nil {
		return err
	}

	return nil
}

// apaga godoc
// @Summary apaga produto por id
// @Tags Produto
// @Produce json
// @Param        id   path      string  true  "id do produto"
// @Router /produto/{id} [delete]
func (h *Produto) apaga(ctx echo.Context) error {
	var (
		produtoID int
		err       error
	)

	id := ctx.Param("id")
	if produtoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}
	p := &domain.Produto{
		Id: int64(produtoID),
	}

	err = h.apagarProduto.Apaga(ctx.Request().Context(), p)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusNoContent, nil)
}

// atualiza godoc
// @Summary atualiza um novo produto
// @Tags Produto
// @Accept json
// @Produce json
// @Router /produto [put]
func (h *Produto) atualiza(ctx echo.Context) error {
	var (
		produto   domain.Produto
		produtoID int
		err       error
	)

	if err = ctx.Bind(&produto); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validateProdutoBody(&produto); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	id := ctx.Param("id")
	if produtoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	err = h.atualizarProduto.Atualiza(ctx.Request().Context(), &produto, int64(produtoID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusOK, nil)
}
