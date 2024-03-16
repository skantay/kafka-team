package usecase

import "github.com/skantay/service/internal/entities"

type msgService interface {
	Produce(msg entities.Entity) error
}

type Usecase struct {
	msgSvc msgService
}

func New(msg msgService) Usecase {
	return Usecase{msg}
}

func (u Usecase) Produce(msg entities.Entity) error {
	return u.msgSvc.Produce(msg)
}
