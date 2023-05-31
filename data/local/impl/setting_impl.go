package impl

import (
	"github.com/jinzhu/gorm"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
)

type settingDao struct {
	db *gorm.DB
}

func NewSettingDao(db *gorm.DB) local.SettingDao {
	return &settingDao{db: db}
}

func (sd *settingDao) Read() (*entity.Setting, error) {
	setting := entity.Setting{}
	err := sd.db.Model(entity.Setting{}).Take(&setting).Error

	if err != nil {
		return &entity.Setting{}, err
	}

	return &setting, nil
}
