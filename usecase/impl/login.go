package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type login struct {
	repo repository.UserRepository
}

func NewLoginUseCase(userRepository repository.UserRepository) usecase.Login {
	return &login{repo: userRepository}
}

func (l *login) Invoke(email string, password string) (*entity.User, error) {
	return l.repo.Login(email, password)
}
