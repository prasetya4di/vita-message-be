package local

import "vita-message-service/data/entity"

type SettingDao interface {
	Read() (*entity.Setting, error)
}
