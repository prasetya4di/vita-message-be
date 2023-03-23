package repository

import "vita-message-service/data/entity"

type EnergyRepository interface {
	Read(email string) (entity.Energy, error)
	Insert(energy entity.Energy) (entity.Energy, error)
	Update(energy entity.Energy) (entity.Energy, error)
}
