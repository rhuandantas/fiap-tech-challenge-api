package handlers

import (
	"fiap-tech-challenge-api/internal/adpters/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -package=mock_handlers -destination=../../../test/mock/handlers/$GOFILE
type Cliente struct {
	repo repository.ClienteRepo
}

func NewCliente(repo repository.ClienteRepo) *Cliente {
	return &Cliente{
		repo: repo,
	}
}

func (h *Cliente) RegistraRotasCliente(server *echo.Echo) {
	server.POST("/cliente", h.insere)
	server.GET("/cliente", h.lista)
}

// insere godoc
// @Summary insere um novo cliente
// @Tags Cliente
// @Accept json
// @Produce json
// @Success 200 {object} domain.Cliente
// @Router /cliente [post]
func (h *Cliente) insere(c echo.Context) error {
	response := make(map[string]string)
	response["status"] = "UP"
	return c.JSON(http.StatusOK, response)
}

// lista godoc
// @Summary Show the status of http.
// @Description get the status of http.
// @Tags Cliente
// @Accept */*
// @Produce json
// @Success 200 {object} domain.Cliente
// @Router /client [get]
func (h *Cliente) lista(c echo.Context) error {
	response := make(map[string]string)
	response["status"] = "OK"
	return c.JSON(http.StatusOK, response)
}
