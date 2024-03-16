package service

import "github.com/skantay/service/internal/entities"

type msgBroker interface {
	Produce(msg entities.Entity) error
}

type Service struct {
	broker msgBroker
}

func New(b msgBroker) *Service {
	return &Service{broker: b}
}

func (s Service) Produce(msg entities.Entity) error {
	return s.broker.Produce(msg)
}
