package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type getUser struct {
	repo repository.UserRepository
}

func NewGetUser(userRepository repository.UserRepository) usecase.GetUser {
	return &getUser{repo: userRepository}
}

func (gu *getUser) Invoke(email string) *entity.User {
	return gu.repo.Get(email)
}
