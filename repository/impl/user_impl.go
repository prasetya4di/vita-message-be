package impl

import (
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
	"vita-message-service/repository"
)

type userRepository struct {
	userDao local.UserDao
}

func NewUserRepository(dao local.UserDao) repository.UserRepository {
	return &userRepository{userDao: dao}
}

func (ur *userRepository) Login(email string, password string) (*entity.User, error) {
	return ur.userDao.Login(email, password)
}

func (ur *userRepository) Register(user entity.User) (*entity.User, error) {
	return ur.userDao.Register(user)
}

func (ur *userRepository) Get(email string) *entity.User {
	return ur.userDao.Get(email)
}
