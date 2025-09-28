package product

import (
	"net/http"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *ProductService
}

func NewProductHandler(service *ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response[[]models.Product]{Data: products})
}
