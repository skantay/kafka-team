package app

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	broker "github.com/skantay/service/internal/infrastructure/kafka"
)

func Run() error {
	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9093",
			"group.id":          "myGroup",

			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		panic(err)
	}

	cons := broker.New(c)

	cons.Consume()

	return nil
}
