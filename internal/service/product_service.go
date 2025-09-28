package service

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}
