package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/models"
	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProducer struct {
	client *kgo.Client
}

func NewKafkaProducer(client *kgo.Client) *KafkaProducer {
	return &KafkaProducer{client: client}
}

func (k KafkaProducer) SendProduct(product models.Product) {
	v, err := json.Marshal(product)

	if err != nil {
		log.Fatal(err)
	}
	record := &kgo.Record{
		Topic:     "products",
		Value:     v,
		Timestamp: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	k.client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		} else {
			fmt.Println("✅ product %s was submitted to kafka sucessfully", product.Name)
		}
	})
}
