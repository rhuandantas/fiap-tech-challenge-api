// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adapters/http"
	"fiap-tech-challenge-api/internal/adapters/http/handlers"
	"fiap-tech-challenge-api/internal/adapters/repository"
	"fiap-tech-challenge-api/internal/core/usecase"
	"fiap-tech-challenge-api/internal/core/usecase/mapper"
	"fiap-tech-challenge-api/internal/util"
)

// Injectors from wire.go:

func InitializeWebServer() (*http.Server, error) {
	healthCheck := handlers.NewHealthCheck()
	dbConnector := repository.NewMySQLConnector()
	clienteRepo := repository.NewClienteRepo(dbConnector)
	cadastrarClienteUseCase := usecase.NewCadastraCliente(clienteRepo)
	pesquisarClientePorCPF := usecase.NewPesquisarClientePorCpf(clienteRepo)
	validator := util.NewCustomValidator()
	cliente := handlers.NewCliente(cadastrarClienteUseCase, pesquisarClientePorCPF, validator)
	produtoRepo := repository.NewProdutoRepo(dbConnector)
	cadastrarProduto := usecase.NewCadastraProduto(produtoRepo)
	pegarProdutoPorCategoria := usecase.NewPegaProdutoPorCategoria(produtoRepo)
	apagarProduto := usecase.NewApagaProduto(produtoRepo)
	atualizarProduto := usecase.NewAtualizaProduto(produtoRepo)
	produto := handlers.NewProduto(validator, cadastrarProduto, pegarProdutoPorCategoria, apagarProduto, atualizarProduto)
	pedidoRepo := repository.NewPedidoRepo(dbConnector)
	pedido := mapper.NewPedidoMapper()
	listarPedidoPorStatus := usecase.NewListaPedidoPorStatus(pedidoRepo, pedido)
	pedidoProdutoRepo := repository.NewPedidoProdutoRepo(dbConnector)
	filaRepo := repository.NewFilaRepo(dbConnector)
	cadastrarPedido := usecase.NewCadastraPedido(pedidoRepo, pedidoProdutoRepo, pedido, clienteRepo, produtoRepo, filaRepo)
	atualizaStatusPedidoUC := usecase.NewAtualizaStatusPedidoUC(pedidoRepo, filaRepo)
	pegarDetalhePedido := usecase.NewPegaDetalhePedido(pedidoRepo, pedidoProdutoRepo, produtoRepo, clienteRepo, pedido)
	realizarCheckout := usecase.NewRealizaCheckout(pedidoRepo, filaRepo)
	handlersPedido := handlers.NewPedido(validator, listarPedidoPorStatus, cadastrarPedido, atualizaStatusPedidoUC, pegarDetalhePedido, realizarCheckout)
	server := http.NewAPIServer(healthCheck, cliente, produto, handlersPedido)
	return server, nil
}
