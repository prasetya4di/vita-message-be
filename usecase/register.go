package usecase

import (
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/request"
)

type Register interface {
	Invoke(request request.RegisterRequest) (*entity.User, error)
}
