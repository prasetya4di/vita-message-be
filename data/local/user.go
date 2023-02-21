package local

import "vita-message-service/data/entity"

type UserDao interface {
	Login(email string, password string) (*entity.User, error)
	Register(user entity.User) (*entity.User, error)
	Get(email string) *entity.User
}
