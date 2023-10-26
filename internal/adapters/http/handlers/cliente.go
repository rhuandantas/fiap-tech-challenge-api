package handlers

import (
	serverErr "fiap-tech-challenge-api/internal/adapters/http/error"
	"fiap-tech-challenge-api/internal/core/commons"
	"fiap-tech-challenge-api/internal/core/domain"
	"fiap-tech-challenge-api/internal/core/usecases"
	"fiap-tech-challenge-api/internal/util"
	"github.com/joomcode/errorx"
	"net/http"
)

type Cliente struct {
	cadastraClienteUC   usecases.CadastrarClienteUseCase
	pegaClientePorCPFUC usecases.PesquisarClientePorCPFUseCase
	validator           util.Validator
}

func NewCliente(cadastraClienteUC usecases.CadastrarClienteUseCase, pegaClientePorCPFUC usecases.PesquisarClientePorCPFUseCase, validator util.Validator) *Cliente {
	return &Cliente{
		cadastraClienteUC:   cadastraClienteUC,
		pegaClientePorCPFUC: pegaClientePorCPFUC,
		validator:           validator,
	}
}

func (h *Cliente) RegistraRotasCliente(server *echo.Echo) {
	server.POST("/cliente", h.cadastra)
	server.GET("/clientes/:cpf", h.pegaPorCpf)
}

// cadastra godoc
// @Summary cadastra um novo cliente
// @Tags Cliente
// @Accept json
// @Produce json
// @Success 200 {object} domain.Cliente
// @Router /cliente [post]
func (h *Cliente) cadastra(ctx echo.Context) error {
	var (
		cliente domain.Cliente
		err     error
	)

	if err = ctx.Bind(&cliente); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	if err = h.validateClienteBody(&cliente); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
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
// @Param        cpf   path      string  true  "cpf do cliente"
// @Success 200 {object} domain.Cliente
// @Router /clientes/{cpf} [get]
func (h *Cliente) pegaPorCpf(ctx echo.Context) error {
	cpf := ctx.Param("cpf")
	c := &domain.Cliente{
		Cpf: cpf,
	}

	if err := c.ValidateCPF(); err != nil {
		return serverErr.HandleError(ctx, commons.BadRequest.New(err.Error()))
	}

	cliente, err := h.pegaClientePorCPFUC.Pesquisa(ctx.Request().Context(), c)
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
