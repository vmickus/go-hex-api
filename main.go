package main

import (
	"github.com/vmickus/go-hexagonal-api/api"
	"github.com/vmickus/go-hexagonal-api/config"
	"github.com/vmickus/go-hexagonal-api/database"
	"github.com/vmickus/go-hexagonal-api/internal/service"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	cfg, err := config.NewApiConfig()
	if err != nil {
		return err
	}

    db := database.NewPostgres()

	postSvc := service.NewPostService(db)

	server := api.NewServer().
		WithPostService(postSvc)

	server.Routes()

	if err := server.Router.Start(cfg.ApiConfig.Host); err != nil {
		return err
	}

	return nil
}
