package impl

import (
	"github.com/gin-gonic/gin"
	"vita-message-service/data/entity"
	"vita-message-service/repository"
	"vita-message-service/usecase"
	"vita-message-service/util/token"
)

type getCurrentUser struct {
	repo repository.UserRepository
}

func NewGetCurrentUser(userRepository repository.UserRepository) usecase.GetCurrentUser {
	return &getCurrentUser{repo: userRepository}
}

func (gcu *getCurrentUser) Invoke(c *gin.Context) (*entity.User, error) {
	email, err := token.ExtractTokenID(c)
	if err != nil {
		return &entity.User{}, err
	}

	user := gcu.repo.Get(email)
	return user, nil
}
