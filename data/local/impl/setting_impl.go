package impl

import (
	"errors"
	"gorm.io/gorm"
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
	err := sd.db.Model(entity.Setting{}).First(&setting).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			setting = entity.Setting{
				SystemContext: "Vita is an AI that help user to answer their question.",
				AiModel:       "gpt-3.5-turbo",
				Temperature:   0.8,
				MaxTokens:     256,
			}
			err = sd.db.Create(&setting).Error
			if err != nil {
				return &entity.Setting{}, err
			}
		} else {
			return &entity.Setting{}, err
		}
	}

	return &setting, nil
}
