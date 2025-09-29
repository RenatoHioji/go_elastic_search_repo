package search

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
)

type SearchService struct {
	repo *SearchRepository
}

func NewSearchService(searchRepository *SearchRepository) *SearchService {
	return &SearchService{repo: searchRepository}
}

func (service SearchService) SearchProduct(q string) (products []models.Product, err error) {
	return service.repo.SearchProducts(q)
}
