package main

import (
	"github.com/vmickus/go-hexagonal-api/internal/adapter/api"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	cfg, err := NewApiConfig()
	if err != nil {
		return err
	}

	server := api.NewServer()
	server.Routes()
	server.LogRoutes()

	if err := server.Router.Start(cfg.ApiConfig.Host); err != nil {
		return err
	}

	return nil
}
