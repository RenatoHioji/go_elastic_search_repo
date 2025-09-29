package config

import "github.com/twmb/franz-go/pkg/kgo"

func InitKafka(cfg Config) (*kgo.Client, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.KafkaUrl),
		kgo.DefaultProduceTopic("products"),
	)
	if err != nil {
		return nil, err
	}

	return cl, nil
}
