//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/messaging"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/product"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/search"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/user"
	"github.com/google/wire"
)

func InitializeApp() (*app.App, error) {

	wire.Build(
		config.LoadConfig,
		config.InitES,
		config.InitDB,
		config.InitKafka,
		messaging.NewKafkaProducer,
		user.UserSet,
		search.SearchSet,
		product.ProductSet,
		wire.Struct(new(app.App), "*"),
	)
	return &app.App{}, nil
}
