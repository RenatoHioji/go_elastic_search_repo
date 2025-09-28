//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/handlers"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/repository"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/service"
	"github.com/google/wire"
)

func InitializeApp() (*app.App, error) {

	wire.Build(
		config.LoadConfig,
		//config.InitES,
		config.InitDB,
		repository.NewProductRepository,
		service.NewProductService,
		handlers.NewProductHandler,
		wire.Struct(new(app.App), "*"),
	)
	return &app.App{}, nil
}
