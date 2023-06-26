package singleton

import (
	"log"
	"sync"

	"veterinary-employee/settings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer
var singletonKafkaClient sync.Once

func initProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": settings.InitializeKafka().BootstrapServers,
		"client.id": settings.InitializeKafka().GroupId,
		"acks": "all",
	})

	if err != nil {
		log.Panic("error while creating a producer: ", err.Error())
	}

	producer = p
}

func NewProducer() *kafka.Producer {
	singletonKafkaClient.Do(initProducer)
	return producer
}
