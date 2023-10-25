// wire.go
//go:build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/adpters/database"
	"fiap-tech-challenge-api/internal/adpters/http"
	"fiap-tech-challenge-api/internal/adpters/http/handlers"
	"fiap-tech-challenge-api/internal/adpters/repository"
	"github.com/google/wire"
)

func InitializeWebServer() (*http.Server, error) {
	wire.Build(database.NewMySQLConnector,
		repository.NewClienteRepo,
		handlers.NewCliente,
		handlers.NewHealthCheck,
		http.NewAPIServer)
	return &http.Server{}, nil
}
