package handlers

import (
	serverErr "fiap-tech-challenge-api/internal/adapters/http/error"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/util"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"

	"github.com/joomcode/errorx"
)

type Pedido struct {
	validator           util.Validator
	listaPorStatusUC    usecase.ListarPedidoPorStatus
	cadastraPedidoUC    usecase.CadastrarPedido
	atualizaStatusUC    usecase.AtualizaStatusPedidoUC
	pegaDetalhePedidoUC usecase.PegarDetalhePedido
	realizaCheckoutUC   usecase.RealizarCheckout
}

func NewPedido(validator util.Validator,
	listaPorStatusUC usecase.ListarPedidoPorStatus,
	cadastraPedidoUC usecase.CadastrarPedido,
	atualizaStatusUC usecase.AtualizaStatusPedidoUC,
	pegaDetalhePedidoUC usecase.PegarDetalhePedido,
	realizaCheckoutUC usecase.RealizarCheckout,
) *Pedido {
	return &Pedido{
		validator:           validator,
		listaPorStatusUC:    listaPorStatusUC,
		cadastraPedidoUC:    cadastraPedidoUC,
		atualizaStatusUC:    atualizaStatusUC,
		pegaDetalhePedidoUC: pegaDetalhePedidoUC,
		realizaCheckoutUC:   realizaCheckoutUC,
	}
}

func (h *Pedido) RegistraRotasPedido(server *echo.Echo) {
	server.POST("/pedido", h.cadastra)
	server.GET("/pedidos/:statuses", h.listaPorStatus)
	server.GET("/pedido/detail/:id", h.listaDetail)
	server.PATCH("/pedido/:id", h.atualizaStatus)
	server.PATCH("/pedido/checkout/:pedidoId", h.checkout)
}

// cadastra godoc
// @Summary cadastra um novo pedido
// @Tags Pedido
// @Accept json
// @Produce json
// @Router /pedido [post]
func (h *Pedido) cadastra(ctx echo.Context) error {
	var (
		req domain.PedidoRequest
		err error
	)

	if err = ctx.Bind(&req); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validatePedidoBody(&req); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	response, err := h.cadastraPedidoUC.Cadastra(ctx.Request().Context(), &req)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusCreated, echo.Map{"id": response.Id})
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

func (h *Pedido) validatePedidoBody(p *domain.PedidoRequest) error {
	if err := h.validator.ValidateStruct(p); err != nil {
		return err
	}

	// check out something more?

	return nil
}

// atualizaStatus godoc
// @Summary atualiza o status do pedido
// @Tags Pedido
// @Accept json
// @Produce json
// @Router /pedido/{id} [patch]
func (h *Pedido) atualizaStatus(ctx echo.Context) error {
	var (
		status struct {
			Status string `json:"status"`
		}
		pedidoID int
		err      error
	)

	if err = ctx.Bind(&status); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	id := ctx.Param("id")
	if pedidoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	err = h.atualizaStatusUC.Atualiza(ctx.Request().Context(), status.Status, int64(pedidoID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusOK, status)
}

// listaDetail godoc
// @Summary lista detalhes do pedido
// @Tags Pedido
// @Produce json
// @Param        id   path      integer  true  "id do pedido a ser lista"
// @Success 200 {object} domain.Pedido
// @Router /pedido/detail/{id} [get]
func (h *Pedido) listaDetail(ctx echo.Context) error {
	var (
		pedidoID int
		err      error
	)

	id := ctx.Param("id")
	if pedidoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	pedido, err := h.pegaDetalhePedidoUC.Pesquisa(ctx.Request().Context(), int64(pedidoID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, pedido)
}

// checkout godoc
// @Summary checkout do pedido
// @Tags Pedido
// @Accept json
// @Param        pedidoId   path      integer  true  "id do pedido a ser feito o checkout"
// @Produce json
// @Router /pedido/checkout/{pedidoId} [patch]
func (h *Pedido) checkout(ctx echo.Context) error {
	var (
		pedidoID int
		err      error
	)

	id := ctx.Param("pedidoId")
	if pedidoID, err = strconv.Atoi(id); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(fmt.Sprintf("%s não é um id válido", id)))
	}

	err = h.realizaCheckoutUC.FakeCheckout(ctx.Request().Context(), int64(pedidoID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.NoContent(http.StatusOK)
}
