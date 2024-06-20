package handlers

import (
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	serverErr "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"net/http"
)

type Login struct {
	pegaClientePorCPFUC usecase.PesquisarCliente
	tokenJwt            auth.Token
}

func NewLogin(pegaClientePorCPFUC usecase.PesquisarCliente, tokenJwt auth.Token) *Login {
	return &Login{
		tokenJwt:            tokenJwt,
		pegaClientePorCPFUC: pegaClientePorCPFUC,
	}
}

func (h *Login) RegistraRotasLogin(server *echo.Echo) {
	server.GET("/login/:cpf", h.login)
}

// login godoc
// @Summary pega um cliente por cpf
// @Tags Login
// @Accept */*
// @Produce json
// @Param        cpf   path      string  true  "cpf do cliente"
// @Router /login/{cpf} [get]
func (h *Login) login(ctx echo.Context) error {
	cpf := ctx.Param("cpf")
	c := &domain.Cliente{
		Cpf: &cpf,
	}

	if err := c.ValidateCPF(); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	cliente, err := h.pegaClientePorCPFUC.PesquisaPorCPF(ctx.Request().Context(), c)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	if cliente == nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New("user not found"))
	}

	token, err := h.tokenJwt.GenerateToken(cpf)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("Authorization", token)

	return ctx.JSON(http.StatusOK, "user authenticated, check the Authorization header")
}
