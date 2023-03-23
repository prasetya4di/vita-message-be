package impl

import (
	"fmt"
	"time"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type addEnergy struct {
	repository repository.EnergyRepository
}

func NewAddEnergy(repository repository.EnergyRepository) usecase.AddEnergy {
	return &addEnergy{repository: repository}
}

func (ae *addEnergy) Invoke(email string) (*entity.Energy, error) {
	newEnergy := entity.Energy{
		Email:       email,
		Energy:      100,
		ExpiredDate: time.Now().AddDate(0, 1, 0).Local(),
	}

	energy, err := ae.repository.Insert(newEnergy)
	if err != nil {
		return &newEnergy, fmt.Errorf("error when insery energy: %v", err)
	}

	return &energy, nil
}
