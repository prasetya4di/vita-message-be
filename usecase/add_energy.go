package usecase

import "vita-message-service/data/entity"

type AddEnergy interface {
	Invoke(email string) (*entity.Energy, error)
}
