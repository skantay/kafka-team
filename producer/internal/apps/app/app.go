package app

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	v1 "github.com/skantay/service/internal/controllers/http/v1"
	broker "github.com/skantay/service/internal/infrastructure/kafka"
	"github.com/skantay/service/internal/service"
	"github.com/skantay/service/internal/usecase"
)

func Run() error {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9093",
		},
	)
	if err != nil {
		panic(err)
	}

	prod := broker.New(p)

	service := service.New(prod)

	usecase := usecase.New(service)

	v1.Run(usecase)

	return nil
}
