package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/usecase"
)

type saveMessage struct {
	dao local.MessageDao
}

func NewSaveMessage(dao local.MessageDao) usecase.SaveMessage {
	return &saveMessage{dao: dao}
}

func (sm *saveMessage) Invoke(message entity.Message) (entity.Message, error) {
	return sm.dao.Insert(message)
}
