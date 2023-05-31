package impl

import (
	"fmt"
	"gorm.io/gorm"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
)

type energyDao struct {
	db *gorm.DB
}

func NewEnergyDao(db *gorm.DB) local.EnergyDao {
	return &energyDao{
		db: db,
	}
}

func (ed *energyDao) Read(email string) (entity.Energy, error) {
	energy := entity.Energy{}

	err := ed.db.Where("email = ?", email).Take(&energy).Error
	if err != nil {
		return entity.Energy{}, fmt.Errorf("error when read energy: %v", err)
	}

	return energy, nil
}

func (ed *energyDao) Insert(energy entity.Energy) (entity.Energy, error) {
	err := ed.db.Create(&energy).Error
	if err != nil {
		return energy, fmt.Errorf("error when insert energy: %v", err)
	}
	return energy, nil
}

func (ed *energyDao) Update(energy entity.Energy) (entity.Energy, error) {
	err := ed.db.Update(energy.TableName(), &energy).Error
	if err != nil {
		return energy, fmt.Errorf("error when update energy: %v", err)
	}
	return energy, nil
}
