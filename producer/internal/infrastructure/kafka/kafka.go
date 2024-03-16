package broker

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/skantay/service/internal/entities"
)

type Producer struct {
	prod *kafka.Producer
}

func New(prod *kafka.Producer) Producer {
	return Producer{prod}
}

func (p Producer) Produce(msg entities.Entity) error {
	topic := "topic-test"

	return p.prod.Produce(

		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("%v", msg)),
		}, nil,
	)
}
