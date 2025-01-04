package kafka

import (
	"strings"

	"github.com/peixotoleonardo/kafka/golang/internal/utils/env"
)

type config struct {
	brokers []string
	topic   string
}

func newConfig() (*config, error) {
	brokers, err := env.GetEnv("KAFKA_BROKERS")

	if err != nil {
		return nil, err
	}

	topic, err := env.GetEnv("KAFKA_TOPIC")

	if err != nil {
		return nil, err
	}

	return &config{topic: topic, brokers: strings.Split(brokers, ",")}, nil
}
