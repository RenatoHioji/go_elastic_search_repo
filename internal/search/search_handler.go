package search

import (
	"net/http"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	service *SearchService
}

func NewSearchHandler(service *SearchService) *SearchHandler {
	return &SearchHandler{service: service}
}

func (h SearchHandler) SearchProduct(c *gin.Context) {
	query := c.Query("q")
	search, err := h.service.SearchProduct(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response[[]models.Product]{Data: search})
}
