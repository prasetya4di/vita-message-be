package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/delivery/rest/request"
	"vita-message-service/repository"
	"vita-message-service/usecase"
)

type register struct {
	repo repository.UserRepository
}

func NewRegisterUseCase(userRepository repository.UserRepository) usecase.Register {
	return &register{
		repo: userRepository,
	}
}

func (r *register) Invoke(request request.RegisterRequest) (*entity.User, error) {
	newUser := entity.User{
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Nickname:  request.Nickname,
		BirthDate: request.BirthDate,
	}

	return r.repo.Register(newUser)
}
