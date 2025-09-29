package search

import (
	"encoding/json"
	"strings"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/elastic/go-elasticsearch/v9"
)

type SearchRepository struct {
	es *elasticsearch.Client
}

func NewSearchRepository(es *elasticsearch.Client) *SearchRepository {
	return &SearchRepository{es: es}
}

func (r SearchRepository) SearchProducts(query string) ([]models.Product, error) {
	res, err := r.es.Search(
		r.es.Search.WithIndex("products"),
		r.es.Search.WithBody(strings.NewReader(`{"query":{"match":{"name":"`+query+`"}}}`)),
		r.es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var esResp struct {
		Hits struct {
			Hits []struct {
				Source models.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&esResp); err != nil {
		return nil, err
	}

	products := make([]models.Product, len(esResp.Hits.Hits))
	for i, hit := range esResp.Hits.Hits {
		products[i] = hit.Source
	}
	return products, nil
}
