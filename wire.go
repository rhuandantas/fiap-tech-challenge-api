// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"fiap-tech-challenge-api/internal/server"
	"fiap-tech-challenge-api/internal/server/handlers"
	"github.com/google/wire"
)

func InitializeWebServer() (*server.HttpServer, error) {
	wire.Build(handlers.NewHealthCheck,
		server.NewAPIServer)
	return &server.HttpServer{}, nil
}
