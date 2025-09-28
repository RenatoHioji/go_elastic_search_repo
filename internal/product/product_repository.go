package product

import (
	"database/sql"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo ProductRepository) GetAllProducts() (products []models.Product, err error) {
	rows, err := repo.db.Query("SELECT id, name FROM products")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
