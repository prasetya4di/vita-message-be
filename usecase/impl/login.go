package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/request"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type login struct {
	repo repository.UserRepository
}

func NewLoginUseCase(userRepository repository.UserRepository) usecase.Login {
	return &login{repo: userRepository}
}

func (l *login) Invoke(request request.LoginRequest) (*entity.User, error) {
	return l.repo.Login(request.Email, request.Password)
}
