package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/model"
	"github.com/elastic/go-elasticsearch/v9"
)

func main() {
	//es, _ := elasticsearch.NewDefaultClient()
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		// Password: password...
		// Username: username...
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	log.Println(es.Info())
	document := model.Document{Name: "go-elastic-search"}

	data, _ := json.Marshal(document)
	es.Index("my_index", bytes.NewReader(data))

	query := `{ "query: { "match_all": {} } }"`

	es.Search(
		es.Search.WithIndex("my_index"),
		es.Search.WithQuery(query),
	)
	es.Update("my_index", "id", strings.NewReader(`{doc: {language: "Go"}}`))
	es.Delete("my_index", "id")

	es.Indices.Delete([]string{"my_index"}) // Deleting index
}
