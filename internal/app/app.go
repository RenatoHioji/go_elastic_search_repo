package app

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/product"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/search"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/user"
)

type App struct {
	ProductHandler *product.ProductHandler
	UserHandler    *user.UserHandler
	SearchHandler  *search.SearchHandler
	Config         config.Config
}
