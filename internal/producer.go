package internal

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ProducerService struct {
	p *kafka.Producer
}

type Config struct {
	bootstrapServers string
	clientId         string
	key              string
	secret           string
}

func NewProducer() (*ProducerService, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "blah:9092",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     "blah",
		"sasl.password":     "blah",
		"security.protocol": "sasl_ssl",
		"client.id":         "go-project",
	})
	if err != nil {
		return nil, err
	}
	return &ProducerService{
		p: p,
	}, nil
}

func (producer *ProducerService) sendMessage(msg string, channel chan kafka.Event, topic string) error {

	return producer.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(msg),
	}, channel)

}
