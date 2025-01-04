package kafka

import (
	"github.com/IBM/sarama"
)

type (
	Producer interface {
		Send(value string) (bool, error)
		Close() error
	}

	producer struct {
		producer sarama.AsyncProducer
		config   *config
	}
)

func NewProducer() (Producer, error) {
	config, err := newConfig()

	if err != nil {
		return nil, err
	}

	p, err := sarama.NewAsyncProducer(config.brokers, sarama.NewConfig())

	if err != nil {
		return nil, err
	}

	return &producer{producer: p, config: config}, nil
}

func (p *producer) Send(value string) (bool, error) {
	message := &sarama.ProducerMessage{
		Topic: p.config.topic,
		Value: sarama.StringEncoder(value),
	}

	p.producer.Input() <- message

	select {
	case err := <-p.producer.Errors():
		return false, err.Err
	default:
		return true, nil
	}
}

func (p *producer) Close() error {
	return p.producer.Close()
}
