// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/util"
	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(repository.NewMySQLConnector,
		util.NewCustomValidator,
		repository.NewClienteRepo,
		repository.NewProdutoRepo,
		usecase.NewCadastraCliente,
		usecase.NewPesquisarClientePorCpf,
		usecase.NewCadastraProduto,
		usecase.NewPegaProdutoPorCategoria,
		usecase.NewApagaProduto,
		usecase.NewAtualizaProduto,
		handlers.NewCliente,
		handlers.NewProduto,
		handlers.NewHealthCheck,
		http.NewAPIServer)
	return &http.Server{}, nil
}
