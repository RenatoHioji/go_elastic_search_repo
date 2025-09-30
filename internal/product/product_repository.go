package product

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r ProductRepository) GetAllProducts() (products []models.Product, err error) {
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductRepository) CreateProduct(product *models.Product) (prod *models.Product, err error) {
	if err := r.db.Create(&product).Error; err != nil {
		log.Fatal(err)
	}

	return product, nil
}
