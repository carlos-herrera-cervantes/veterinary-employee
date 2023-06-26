package services

import (
	"veterinary-employee/singleton"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaService struct {
	Producer singleton.IKafka
}

func (k *KafkaService) SendMessage(topic string, message []byte) error {
	if err := k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}, nil); err != nil {
		return err
	}

	return nil
}
