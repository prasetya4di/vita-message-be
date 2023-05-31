package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/repository"
)

type settingRepository struct {
	settingDao local.SettingDao
}

func NewSettingRepository(settingDao local.SettingDao) repository.SettingRepository {
	return &settingRepository{settingDao: settingDao}
}

func (sr *settingRepository) Read() (*entity.Setting, error) {
	return sr.settingDao.Read()
}
