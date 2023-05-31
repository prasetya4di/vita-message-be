package repository

import "vita-message-service/data/entity"

type SettingRepository interface {
	Read() (*entity.Setting, error)
}
