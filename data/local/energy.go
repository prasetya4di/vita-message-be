package local

import "vita-message-service/data/entity"

type EnergyDao interface {
	Read(email string) (entity.Energy, error)
	Insert(energy entity.Energy) (entity.Energy, error)
	Update(energy entity.Energy) (entity.Energy, error)
}
