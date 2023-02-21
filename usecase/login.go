package usecase

import (
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/request"
)

type Login interface {
	Invoke(request request.LoginRequest) (*entity.User, error)
}
