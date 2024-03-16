package broker

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	cons *kafka.Consumer
}

func New(cons *kafka.Consumer) Consumer {
	return Consumer{cons}
}

func (c Consumer) Consume() {
	topic := "topic-test"

	c.cons.Subscribe(topic, nil)

	for {
		msg, err := c.cons.ReadMessage(time.Second * 2)
		if err == nil {
			fmt.Printf("Got message %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
