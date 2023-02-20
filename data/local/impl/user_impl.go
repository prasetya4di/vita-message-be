package impl

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"vita-message-service/data/entity"
	"vita-message-service/data/local"
)

type userDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) local.UserDao {
	return &userDao{
		db: db,
	}
}

func (ud *userDao) Login(email string, password string) (entity.User, error) {
	// Retrieve user from database
	row := ud.db.QueryRow("SELECT * FROM user WHERE email = ?", email)
	var user entity.User
	err := row.Scan(&user.Email, &user.FirstName, &user.LastName, &user.BirthDate, &user.Nickname, &user.Password)
	if err != nil {
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.User{}, err
	}
	user.Password = ""

	return user, nil
}

func (ud *userDao) Register(user entity.User) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	// Insert user into database
	_, err = ud.db.Exec("INSERT INTO user (email, password, first_name, last_name, nickname, birth_date) VALUES (?, ?, ?, ?, ?, ?)",
		user.Email,
		string(hash),
		user.FirstName,
		user.LastName,
		user.Nickname,
		user.BirthDate)
	if err != nil {
		return entity.User{}, err
	}

	user.Password = ""
	return user, nil
}

func (ud *userDao) Get(email string) entity.User {
	// Retrieve user from database
	row := ud.db.QueryRow("SELECT * FROM user WHERE email = ?", email)
	var user entity.User
	err := row.Scan(&user.Email, &user.FirstName, &user.LastName, &user.BirthDate, &user.Nickname, &user.Password)
	if err != nil {
		return entity.User{}
	}

	return user
}
