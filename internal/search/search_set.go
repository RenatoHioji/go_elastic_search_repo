package search

import "github.com/google/wire"

var SearchSet = wire.NewSet(
	NewSearchHandler,
	NewSearchService,
	NewSearchRepository,
)
