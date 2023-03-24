package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/repository"
)

type energyRepository struct {
	dao local.EnergyDao
}

func NewEnergyRepository(dao local.EnergyDao) repository.EnergyRepository {
	return &energyRepository{dao: dao}
}

func (er *energyRepository) Read(email string) (entity.Energy, error) {
	return er.dao.Read(email)
}

func (er *energyRepository) Insert(energy entity.Energy) (entity.Energy, error) {
	return er.dao.Insert(energy)
}

func (er *energyRepository) Update(energy entity.Energy) (entity.Energy, error) {
	return er.dao.Update(energy)
}
