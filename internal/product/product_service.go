package product

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/messaging"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
)

type ProductService struct {
	repo     *ProductRepository
	producer *messaging.KafkaProducer
}

func NewProductService(repo *ProductRepository, producer *messaging.KafkaProducer) *ProductService {
	return &ProductService{repo: repo, producer: producer}
}

func (s ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s ProductService) CreateProduct(product *models.Product) error {
	prod, err := s.repo.CreateProduct(product)
	if err != nil {
		log.Println("Error creating product:", err)
		return err
	}

	s.producer.SendProduct(prod)

	return nil
}
