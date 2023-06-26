package settings

import "os"

type kafka struct {
	BootstrapServers string
	GroupId string
	Topics topic
}

type topic struct {
	ProfileUpdate string
}

var singletonKafka *kafka

func InitializeKafka() *kafka {
	if singletonKafka != nil {
		return singletonKafka
	}

	lock.Lock()
	defer lock.Unlock()

	singletonKafka = &kafka{
		BootstrapServers: os.Getenv("KAFKA_SERVERS"),
		GroupId: os.Getenv("KAFKA_GROUP_ID"),
		Topics: topic{
			ProfileUpdate: "veterinary-employee-profile-update",
		},
	}

	return singletonKafka
}
