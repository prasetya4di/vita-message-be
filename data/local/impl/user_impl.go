package impl

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
)

type userDao struct {
	db     *sql.DB
	gormDb *gorm.DB
}

func NewUserDao(db *sql.DB, db2 *gorm.DB) local.UserDao {
	return &userDao{
		db:     db,
		gormDb: db2,
	}
}

func (ud *userDao) Login(email string, password string) (*entity.User, error) {
	// Retrieve user from database
	user := entity.User{}
	err := ud.gormDb.Model(entity.User{}).Where("email = ?", email).Take(&user).Error

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &entity.User{}, err
	}
	user.Password = ""

	return &user, nil
}

func (ud *userDao) Register(user entity.User) (*entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &entity.User{}, err
	}

	user.Password = string(hash)
	err = ud.gormDb.Create(&user).Error
	if err != nil {
		return &entity.User{}, err
	}

	user.Password = ""
	return &user, nil
}

func (ud *userDao) Get(email string) *entity.User {
	// Retrieve user from database
	user := entity.User{}
	err := ud.gormDb.Model(entity.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		log.Printf("error when get user :%v", err)
	}

	user.Password = ""
	return &user
}
