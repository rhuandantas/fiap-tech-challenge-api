package handlers

import (
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecase"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	serverErr "github.com/rhuandantas/fiap-tech-challenge-commons/pkg/errors"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/middlewares/auth"
	"github.com/rhuandantas/fiap-tech-challenge-commons/pkg/util"
	"net/http"
	"strconv"
)

type Cliente struct {
	cadastraClienteUC usecase.CadastrarClienteUseCase
	pegaClienteUC     usecase.PesquisarCliente
	lgpd              usecase.LGPD
	validator         util.Validator
	tokenJwt          auth.Token
}

func NewCliente(cadastraClienteUC usecase.CadastrarClienteUseCase, pegaClientePorCPFUC usecase.PesquisarCliente, lgpd usecase.LGPD, validator util.Validator, tokenJwt auth.Token) *Cliente {
	return &Cliente{
		cadastraClienteUC: cadastraClienteUC,
		pegaClienteUC:     pegaClientePorCPFUC,
		lgpd:              lgpd,
		validator:         validator,
		tokenJwt:          tokenJwt,
	}
}

func (h *Cliente) RegistraRotasCliente(server *echo.Echo) {
	server.POST("/cliente", h.cadastra)
	server.GET("/clientes/:cpf", h.pegaPorCpf, h.tokenJwt.VerifyToken)
	server.GET("/internal/clientes/:id", h.pegaPorID)
	server.DELETE("/lgpd/clientes/delete", h.anonimizarClientLGPD)
}

// cadastra godoc
// @Summary cadastra um novo cliente
// @Tags Cliente
// @Accept json
// @Produce json
// @Param			pedido	body		domain.ClienteRequest	true	"cria novo cliente"
// @Success 200 {object} domain.Cliente
// @Router /cliente [post]
func (h *Cliente) cadastra(ctx echo.Context) error {
	var (
		cliente domain.Cliente
		err     error
	)

	if err = ctx.Bind(&cliente); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	if err = h.validateClienteBody(&cliente); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	newCliente, err := h.cadastraClienteUC.Cadastra(ctx.Request().Context(), &cliente)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusCreated, newCliente)
}

// pegaPorCpf godoc
// @Summary pega um cliente por cpf
// @Tags Cliente
// @Accept */*
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        cpf   path      string  true  "cpf do cliente"
// @Success 200 {object} domain.Cliente
// @Router /clientes/{cpf} [get]
func (h *Cliente) pegaPorCpf(ctx echo.Context) error {
	cpf := ctx.Param("cpf")
	
	c := &domain.Cliente{
		Cpf: &cpf,
	}

	if err := c.ValidateCPF(); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	cliente, err := h.pegaClienteUC.PesquisaPorCPF(ctx.Request().Context(), c)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	token, err := h.tokenJwt.GenerateToken(cpf)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("Authorization", token)

	return ctx.JSON(http.StatusOK, cliente)
}

// pegaPorID godoc
// @Summary pega um cliente por id
// @Tags Cliente
// @Accept */*
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param        id   path      string  true  "id do cliente"
// @Success 200 {object} domain.Cliente
// @Router /clientes/{id} [get]
func (h *Cliente) pegaPorID(ctx echo.Context) error {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return serverErr.BadRequest.New(err.Error())
	}

	cliente, err := h.pegaClienteUC.PesquisaPorID(ctx.Request().Context(), int64(intID))
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}

	return ctx.JSON(http.StatusOK, cliente)
}

func (h *Cliente) validateClienteBody(c *domain.Cliente) error {
	if err := h.validator.ValidateStruct(c); err != nil {
		return err
	}

	if err := c.ValidateCPF(); err != nil {
		return err
	}

	return nil
}

// delete godoc
// @Summary deleta um cliente
// @Tags Cliente
// @Accept json
// @Produce json
// @Param			pedido	body		domain.LGPDClienteRequest	true	"anonimiza os dados do cliente"
// @Success 200 {object} domain.Cliente
// @Router /lgpd/clientes/delete [delete]
func (h *Cliente) anonimizarClientLGPD(ctx echo.Context) error {
	var (
		cliente domain.Cliente
		err     error
	)
	
	if err = ctx.Bind(&cliente); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	if err := cliente.ValidateCPF(); err != nil {
		return serverErr.HandleError(ctx, serverErr.BadRequest.New(err.Error()))
	}

	err = h.lgpd.Anonimizar(ctx.Request().Context(), &cliente)
	if err != nil {
		return serverErr.HandleError(ctx, errorx.Cast(err))
	}
	
	return ctx.NoContent(http.StatusOK)

}
