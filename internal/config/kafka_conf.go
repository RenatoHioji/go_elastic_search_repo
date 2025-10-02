package config

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

func InitKafka(cfg Config) (*kgo.Client, error) {
	seeds := []string{cfg.KafkaUrl}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
	)
	if err != nil {
		return nil, err
	}

	return cl, nil
}
