package messaging

import (
	"context"
	"encoding/json"
	"fmt"
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

func (k *KafkaProducer) SendProduct(product *models.Product) {
	v, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("failed to marshal product: %v\n", err)
		return
	}

	record := &kgo.Record{
		Topic:     "products",
		Value:     v,
		Timestamp: time.Now(),
	}

	k.client.Produce(context.Background(), record, func(_ *kgo.Record, err error) {
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		} else {
			fmt.Printf("product %s was submitted to Kafka successfully\n", product.Name)
		}
	})
}
