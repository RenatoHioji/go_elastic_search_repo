package product

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
)

type ProductService struct {
	repo *ProductRepository
}

func NewProductService(repo *ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}
