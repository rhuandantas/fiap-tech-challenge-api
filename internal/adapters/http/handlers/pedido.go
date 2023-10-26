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
	"strings"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type Pedido struct {
	validator        util.Validator
	listaPorStatusUC usecase.ListarPedidoPorStatus
}

func NewPedido(validator util.Validator,
	listaPorStatusUC usecase.ListarPedidoPorStatus,
) *Pedido {
	return &Pedido{
		validator:        validator,
		listaPorStatusUC: listaPorStatusUC,
	}
}

func (h *Pedido) RegistraRotasPedido(server *echo.Echo) {
	server.POST("/pedido", h.cadastra)
	server.GET("/pedidos/:statuses", h.listaPorStatus)
	server.PATCH("/pedido", h.atualizaStatus)
}

// cadastra godoc
// @Summary cadastra um novo pedido
// @Tags Pedido
// @Accept json
// @Produce json
// @Router /pedido [post]
func (h *Pedido) cadastra(ctx echo.Context) error {
	var (
		pedido domain.Pedido
		err    error
	)

	if err = ctx.Bind(&pedido); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validatePedidoBody(&pedido); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	/*newProduto, err := h.cadastraProdutoUC.Cadastra(ctx.Request().Context(), &pedido)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	*/
	return ctx.JSON(http.StatusCreated, nil)
}

// listaPorStatus godoc
// @Summary lista pedido por status
// @Tags Pedido
// @Produce json
// @Param        statuses   path      string  true  "status dos pedidos a ser pesquisado"
// @Success 200 {array} domain.Pedido
// @Router /pedidos/{statuses} [get]
func (h *Pedido) listaPorStatus(ctx echo.Context) error {
	statuses := ctx.Param("statuses")
	filter := strings.Split(statuses, ",")

	pedidos, err := h.listaPorStatusUC.ListaPorStatus(ctx.Request().Context(), filter)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedidos)
}

func (h *Pedido) validatePedidoBody(p *domain.Pedido) error {
	if err := h.validator.ValidateStruct(p); err != nil {
		return err
	}

	// check out something more?

	return nil
}

// atualiza godoc
// @Summary atualiza um pedido
// @Tags Pedido
// @Accept json
// @Produce json
// @Router /pedido [patch]
func (h *Pedido) atualizaStatus(ctx echo.Context) error {
	var (
		pedido   domain.Pedido
		pedidoID int
		err      error
	)

	if err = ctx.Bind(&pedido); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	id := ctx.Param("id")
	if pedidoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	return ctx.JSON(http.StatusOK, pedidoID)
}
