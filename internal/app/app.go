package app

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/handlers"
)

type App struct {
	ProductHandler *handlers.ProductHandler
	UserHandler    *handlers.UserHandler
	Config         config.Config
}
