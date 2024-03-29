package handlers

import (
	serverErr "fiap-tech-challenge-api/internal/adapters/http/error"
	"fiap-tech-challenge-api/internal/adapters/http/middlewares/auth"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/util"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Pagamento struct {
	pesquisaPorPedidoID usecase.PesquisaPagamento
	validator           util.Validator
	tokenJwt            auth.Token
}

func NewPagamento(pesquisaPorPedidoID usecase.PesquisaPagamento, validator util.Validator,
	tokenJwt auth.Token) *Pagamento {
	return &Pagamento{
		pesquisaPorPedidoID: pesquisaPorPedidoID,
		validator:           validator,
		tokenJwt:            tokenJwt,
	}
}

func (h *Pagamento) RegistraRotasPagamento(server *echo.Echo) {
	server.GET("/pagamento/:pedido_id", h.pesquisaPorPedidoId, h.tokenJwt.VerifyToken)
}

// pesquisaPorPedidoId godoc
// @Summary pega um pagamento por pedido id
// @Tags Pagamento
// @Accept */*
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        pedido_id   path      string  true  "id do pedido"
// @Success 200 {object} domain.Pagamento
// @Router /pagamento/{pedido_id} [get]
func (h *Pagamento) pesquisaPorPedidoId(ctx echo.Context) error {
	pedidoID := ctx.Param("pedido_id")

	atoi, err := strconv.Atoi(pedidoID)
	if err != nil {
		return err
	}

	cliente, err := h.pesquisaPorPedidoID.PesquisaPorPedidoID(ctx.Request().Context(), int64(atoi))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	return ctx.JSON(http.StatusOK, cliente)
}
