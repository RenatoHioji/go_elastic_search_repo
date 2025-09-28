package product

import "github.com/google/wire"

var ProductSet = wire.NewSet(
	NewProductRepository,
	NewProductService,
	NewProductHandler,
)
